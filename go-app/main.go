package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Define a Prometheus counter
var requestCounter = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Number of requests received",
	},
	[]string{"method", "endpoint"},
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	host, _ := os.Hostname()
	log.Printf("Incoming request %v\n", r.Header)
	requestCounter.WithLabelValues(r.Method, r.URL.Path).Inc()

	fmt.Fprintf(w, "Hello World! I'm served from %s", host)
}

func main() {

	prometheus.MustRegister(requestCounter)

	mux := http.NewServeMux()

	// Expose the registered metrics via HTTP.
	mux.Handle("/metrics", promhttp.Handler())

	mux.HandleFunc("/hello", helloHandler)

	port := "8080"
	log.Printf("Server is running on port %s\n", port)

	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, mux))
}
