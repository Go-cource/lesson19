package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
)

var httpRequestTotal = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total numbers of HTTP Requests",
	},
	[]string{"path"},
)

func IndexPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Index Page!"))
}

func main() {
	http.HandleFunc("/", IndexPage)
	http.ListenAndServe(":8080", nil)
}
