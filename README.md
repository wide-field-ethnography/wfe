# Wide-Field Ethnography 

[![Build Status](https://travis-ci.org/citwild/wfe.svg?branch=master)](https://travis-ci.org/citwild/wfe)
[![Go Report Card](https://goreportcard.com/badge/github.com/citwild/wfe)](https://goreportcard.com/report/github.com/citwild/wfe)
[![GoDoc](https://godoc.org/github.com/citwild/wfe?status.svg)](https://godoc.org/github.com/citwild/wfe)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

[Wide-field ethnography](http://depts.washington.edu/citw/wordpress/?page_id=55) web app. Currently a work-in-progress.

## Requirements

- [Go](https://golang.org/doc/install)

Make sure you [set up your `GOPATH`](https://golang.org/doc/code.html#GOPATH). 

## Install

```
go get github.com/citwild/wfe/cmd/wfe
```

## Use

```
wfe -h
```

Or, if the `GOPATH`/bin directory is not in your `PATH`, on macOS and Linux:

```
"$GOPATH"/bin/wfe -h
```

On Windows:

```
"%GOPATH%"\bin\wfe -h
```

## Acknowledgements

- [Google I/O talk: Building Sourcegraph, a large-scale code search & cross-reference engine in Go](https://text.sourcegraph.com/google-i-o-talk-building-sourcegraph-a-large-scale-code-search-cross-reference-engine-in-go-1f911b78a82e#.49ahhw3cp)
- <strike>[sourcegraph](https://github.com/sourcegraph/sourcegraph)</strike> (no longer available)
- [thesrc](https://github.com/sourcegraph/thesrc)
- [cadvisor](https://github.com/google/cadvisor)
 
## License

The WFE web app is licensed under the [MIT License](https://opensource.org/licenses/MIT), which is an [open source license](https://opensource.org/docs/osd).