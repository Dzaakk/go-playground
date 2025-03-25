package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	registry = prometheus.NewRegistry()

	httpRequestCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total number of HTTP requests",
	},
	)

	httpDuration = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "HTTP request duration in seconds",
			Buckets: []float64{0.1, 0.2, 0.3, 0.4, 0.5, 1, 2, 5},
		},
	)

	httpActiveRequests = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "http_active_requests",
			Help: "Current number of active HTTP requests",
		},
	)
)

func init() {
	registry.MustRegister(httpDuration)
	registry.MustRegister(httpActiveRequests)

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	timer := prometheus.NewTimer(httpDuration)
	defer timer.ObserveDuration()

	httpRequestCounter.Inc()

	httpActiveRequests.Inc()
	defer httpActiveRequests.Dec()

	time.Sleep(time.Duration(100+rand.Intn(400)) * time.Millisecond)

	w.Write([]byte("Hello, Prometheus Monitoring!"))
}

func main() {
	rand.Seed(time.Now().UnixNano())

	http.HandleFunc("/", helloHandler)
	http.Handle("/metrics", promhttp.HandlerFor(
		registry,
		promhttp.HandlerOpts{},
	))

	log.Println("Server starting on :8080")
	log.Println("Metrics available at /metrics")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
