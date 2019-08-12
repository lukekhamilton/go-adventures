package main

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var fooCount = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "foo_total",
	Help: "Number of foo successfully processed.",
})

func foo() {
	for {
		time.Sleep(time.Second)
		fooCount.Add(1)
	}
}

func main() {
	prometheus.MustRegister(fooCount)

	go foo()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":9090", nil)

}
