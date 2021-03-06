package cli

import (
	"crypto/tls"
	"net/http"
	"strings"

	"net"

	"github.com/citwild/wfe/api"
	"github.com/citwild/wfe/app"
	"github.com/citwild/wfe/app/assets"
	"github.com/citwild/wfe/cli/internal/middleware"
	"github.com/citwild/wfe/service"
	"github.com/citwild/wfe/store/mongostore"
	"github.com/gorilla/mux"
	"github.com/mwitkow/go-grpc-middleware"
	"github.com/soheilhy/cmux"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"gopkg.in/mgo.v2"
	"sourcegraph.com/sourcegraph/go-flags"
)

func init() {
	cmdInits = append(cmdInits, func(parser *flags.Parser) {
		_, err := parser.AddCommand("serve",
			"start web server",
			"Starts an HTTP server running the app and API.",
			&ServeCmd{})
		if err != nil {
			zap.S().Fatal("Failed to add serve command.", err)
		}
	})
}

type ServeCmd struct {
	HTTPAddr  string `long:"http-addr" default:":8080" description:"HTTP listen address for app and gRPC API" env:"WFE_HTTP_ADDR"`
	HTTPSAddr string `long:"https-addr" default:":8443" description:"HTTPS (TLS) listen address for app and gRPC API" env:"WFE_HTTPS_ADDR"`

	CertFile string `long:"tls-cert" description:"certificate file for TLS" env:"WFE_TLS_CERT"`
	KeyFile  string `long:"tls-key" description:"key file for TLS" env:"WFE_TLS_KEY"`

	MGOHost string `long:"mgo-host" default:"localhost" description:"MongoDB host machine where the mongod or mongos is running" env:"WFE_MGO_HOST"`
}

func (c *ServeCmd) Execute(_ []string) error {
	var lvl zap.AtomicLevel
	err := lvl.UnmarshalText([]byte(globalOpt.LogLevel))
	if err != nil {
		return err
	}
	config := zap.NewProductionConfig()
	config.Level = lvl
	log, err := config.Build()
	if err != nil {
		return err
	}
	zap.ReplaceGlobals(log)

	app.Init()

	// database
	session, err := mgo.Dial(c.MGOHost)
	if err != nil {
		return err
	}
	defer session.Close()

	// gRPC API
	grpcConfig := &grpcConfig{}
	grpcConfig.servers = service.NewServers()
	grpcConfig.opts = []grpc.ServerOption{
		grpc_middleware.WithUnaryServerChain(
			middleware.NewUnaryServiceInjector(grpcConfig.servers),
			middleware.NewUnaryStoreInjector(mongostore.NewStores(session))),
	}

	// web app
	httpHandler := http.NewServeMux()
	httpHandler.Handle("/", app.NewHandler(app.NewRouter(mux.NewRouter())))
	httpHandler.Handle("/assets/", http.StripPrefix("/assets", assets.NewHandler(mux.NewRouter())))

	err = serveHTTP(c.HTTPAddr, grpcConfig, httpHandler)
	if err != nil {
		return err
	}

	useTLS := c.CertFile != "" || c.KeyFile != ""
	if useTLS {
		err = serveHTTPS(c.HTTPSAddr, grpcConfig, httpHandler, c.CertFile, c.KeyFile)
		if err != nil {
			return err
		}
	}

	select {}
}

func serveHTTP(addr string, grpcConfig *grpcConfig, httpHandler http.Handler) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	// main multiplexer
	mx := cmux.New(lis)

	// gRPC API
	grpcSrv := api.NewServer(grpcConfig.servers, grpcConfig.opts...)

	// web app
	httpSrv := &http.Server{}
	httpSrv.Addr = addr
	httpSrv.Handler = httpHandler

	// multiplex connection between gRPC API and app
	grpcLis := mx.Match(cmux.HTTP2HeaderField("content-type", "application/grpc"))
	httpLis := mx.Match(cmux.Any())

	zap.S().Info("HTTP running.", "addr", addr)
	go func() {
		zap.S().Fatal("Failed to start grpc server.", grpcSrv.Serve(grpcLis))
	}()
	go func() {
		zap.S().Fatal("Failed to start http server.", httpSrv.Serve(httpLis))
	}()
	go func() {
		zap.S().Fatal("Failed to start main multiplexer.", mx.Serve())
	}()

	return nil
}

func serveHTTPS(addr string, grpcConfig *grpcConfig, httpHandler http.Handler, certFile string, keyFile string) error {
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return err
	}

	config := &tls.Config{Certificates: []tls.Certificate{cert}}

	lis, err := tls.Listen("tcp", addr, config)
	if err != nil {
		return err
	}

	// main server
	srv := &http.Server{}
	srv.TLSConfig = config

	// gRPC API
	grpcSrv := api.NewServer(grpcConfig.servers, grpcConfig.opts...)

	// multiplex connection between gRPC API and app
	srv.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcSrv.ServeHTTP(w, r)
		} else {
			httpHandler.ServeHTTP(w, r)
		}
	})

	zap.S().Info("HTTPS running.", "addr", addr)
	go func() {
		zap.S().Fatal("Failed to start https server.", srv.Serve(lis))
	}()

	return nil
}

type grpcConfig struct {
	servers api.Servers
	opts    []grpc.ServerOption
}
