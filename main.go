package main

import (
	"fmt"
	"net/http/httputil"
	"net/url"
	"os"
	"load_balancer/src/handleErr"
)

type simpleServer struct {
	addr  string
	proxy *httputil.ReverseProxy 
}

func newSimpleServer(addr string) *simpleServer {
	serverURL, err := url.Parse(addr)
	handleErr(err)

	return &simpleServer{
		addr:  addr,
		proxy: httputil.NewSingleHostReverseProxy(serverURL),
	}
}

func handleErr(err error) {
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
}


