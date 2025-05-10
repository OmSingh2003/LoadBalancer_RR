package src

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type ReverseProxy struct {
	backendURL *url.URL
	proxy      *httputil.ReverseProxy
}

func NewReverseProxy(backendURL string) (*ReverseProxy, error) {
	serverURL, err := url.Parse(backendURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse backend URL: %w", err)
	}
	return &ReverseProxy{
		backendURL: serverURL,
		proxy:      httputil.NewSingleHostReverseProxy(serverURL),
	}, nil
}

func (rp *ReverseProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rp.proxy.ServeHTTP(w, r)
}}
