package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var httpRequestTotal = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total numbers of HTTP Requests",
	},
	[]string{"path"},
)

func IndexPage(w http.ResponseWriter, r *http.Request) {
	httpRequestTotal.WithLabelValues(r.URL.Path).Inc()
	w.Write([]byte("Index Page!"))
}

func main() {
	prometheus.MustRegister(httpRequestTotal)

	http.HandleFunc("/", IndexPage)
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}
