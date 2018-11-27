package main

import (
	"log"
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

func logSetup() {
	aCondtionUrl := os.Getenv("A_CONDITION_URL")
	bCondtionUrl := os.Getenv("B_CONDITION_URL")
	defaultCondtionUrl := os.Getenv("DEFAULT_CONDITION_URL")

	log.Printf("Server will run on: %s\n", getListenAddress())
	log.Printf("Redirecting to A url: %s\n", aCondtionUrl)
	log.Printf("Redirecting to B url: %s\n", bCondtionUrl)
	log.Printf("Redirecting to Default url: %s\n", defaultCondtionUrl)
}

func handleRequestAndRedirect(res http.ResponseWriter, req *http.Request) {

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
	logSetup()

	http.HandleFunc("/", handleRequestAndRedirect)
	if err := http.ListenAndServe(getListenAddress(), nil); err != nil {
		panic(err)
	}
}
