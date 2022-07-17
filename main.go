package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	port     = kingpin.Flag("port", "Port to start the metrics server on").Short('p').Default("9439").String()
	bindHost = kingpin.Flag("host", "IP address to start the metrics server on").Default("0.0.0.0").String()
)

func main() {
	kingpin.HelpFlag.Short('h')
	kingpin.Parse()

	recorder := NewRecorder()
	collector := NewShellyHTCollector(recorder)

	prometheus.MustRegister(collector)

	log.Printf("Attaching /metrics handler")
	http.HandleFunc("/report", recorder.ServeHTTP)
	http.Handle("/metrics", promhttp.Handler())
	log.Printf("Starting exporter on %s:%s", *bindHost, *port)
	log.Fatal(http.ListenAndServe(*bindHost+":"+*port, nil))
}
