package server

import (
	"github.com/citwild/wfe/api"
	"github.com/citwild/wfe/services"
	"github.com/citwild/wfe/services/internal/localstores"
	"github.com/citwild/wfe/services/internal/middleware"
	"google.golang.org/grpc"
)

func New() *grpc.Server {
	svcs := services.NewServices()
	strs := localstores.NewLocalStores()
	i := middleware.NewInjector(svcs, strs)
	s := grpc.NewServer(grpc.UnaryInterceptor(i.Inject))
	RegisterAll(s, svcs)
	return s
}

func RegisterAll(s *grpc.Server, svcs services.Services) {
	if svcs.Accounts != nil {
		api.RegisterAccountsServer(s, svcs.Accounts)
	}
}
