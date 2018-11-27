package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func serverReverseProxy(target string, res http.ResponseWriter, req *http.Request) {
	//parse the url
	url, _ := url.Parse(target)

	proxy := httputil.NewSingleHostReverseProxy(url)

	req.URL.Host = url.Host
	req.URL.Scheme = url.Scheme
	req.Header.Set("X-forward-Host", req.Header.Get("Host"))
	req.Host = url.Host

	proxy.ServeHTTP(res, req)
}

func main() {

}
