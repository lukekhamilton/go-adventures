// Ref: https://sysdig.com/blog/prometheus-metrics/
package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/common/log"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	counter = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "golang",
		Name:      "my_counter",
		Help:      "This is my counter",
	})

	gauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "golang",
		Name:      "my_gauge",
		Help:      "This is my gauge",
	})

	histogram = prometheus.NewHistogram(prometheus.HistogramOpts{
		Namespace: "golang",
		Name:      "my_histogram",
		Help:      "This is my histogram",
	})

	summary = prometheus.NewSummary(prometheus.SummaryOpts{
		Namespace: "golang",
		Name:      "my_summary",
		Help:      "This is my summary",
	})
)

func main() {
	fmt.Println("booting...")
	rand.Seed(time.Now().Unix())

	prometheus.MustRegister(counter)
	prometheus.MustRegister(gauge)
	prometheus.MustRegister(histogram)
	prometheus.MustRegister(summary)

	http.Handle("/metrics", promhttp.Handler())

	go func() {
		for {
			fmt.Println("Generating stats...")
			counter.Add(1)
			gauge.Add(1)
			histogram.Observe(2)
			summary.Observe(2)

			time.Sleep(time.Second)
		}
	}()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
