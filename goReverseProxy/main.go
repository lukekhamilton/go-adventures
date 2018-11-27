package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

func getListenAddress() string {
	port := getEnv("PORT", "1338")
	return ":" + port
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

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
