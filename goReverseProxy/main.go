package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
)

type requestPayloadStruct struct {
	ProxyCondition string `json:"proxy_condtion"`
}

func parseRequestBody(req *http.Request) requestPayloadStruct {
	decoder := requestBodyDecoder(req)
	var requestPayload requestPayloadStruct
	err := decoder.Decode(&requestPayload)
	if err != nil {
		panic(err)
	}
	log.Printf("requestPayload: %s", requestPayload)
	return requestPayload
}

func requestBodyDecoder(req *http.Request) *json.Decoder {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Panicf("Error reading body: %v", err)
		panic(err)
	}

	req.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	return json.NewDecoder(ioutil.NopCloser(bytes.NewBuffer(body)))
}

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
	aCondtionURL := os.Getenv("A_CONDITION_URL")
	bCondtionURL := os.Getenv("B_CONDITION_URL")
	defaultCondtionURL := os.Getenv("DEFAULT_CONDITION_URL")

	log.Printf("Server will run on: %s\n", getListenAddress())
	log.Printf("Redirecting to A url: %s\n", aCondtionURL)
	log.Printf("Redirecting to B url: %s\n", bCondtionURL)
	log.Printf("Redirecting to Default url: %s\n", defaultCondtionURL)
}

func logRequestPayload(requestionPayload requestPayloadStruct, proxyUrl string) {
	log.Printf("proxy_condition: %s, proxy_url: %s\n", requestionPayload.ProxyCondition, proxyUrl)
}

func getProxyURL(proxyConditionRaw string) string {
	proxyCondition := strings.ToUpper(proxyConditionRaw)
	a_condtion_url := os.Getenv("A_CONDITION_URL")
	b_condtion_url := os.Getenv("B_CONDITION_URL")
	default_condtion_url := os.Getenv("DEFAULT_CONDITION_URL")

	if proxyCondition == "A" {
		return a_condtion_url
	}

	if proxyCondition == "B" {
		return b_condtion_url
	}

	return default_condtion_url
}

func handleRequestAndRedirect(res http.ResponseWriter, req *http.Request) {
	reqPayLoad := parseRequestBody(req)
	url := getProxyURL(reqPayLoad.ProxyCondition)
	logRequestPayload(reqPayLoad, url)
	serveReverseProxy(url, res, req)
}

func serveReverseProxy(target string, res http.ResponseWriter, req *http.Request) {
	//parse the url
	url, _ := url.Parse(target)

	// Create reverse proxy
	proxy := httputil.NewSingleHostReverseProxy(url)

	// Update headers to allow for SSL redirection
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
