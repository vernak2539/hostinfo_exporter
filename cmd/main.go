package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/vernak2539/hostinfo_exporter/src/services"
)

var address = flag.String("listen-address", ":8765", "The address to listen for HTTP requests.")
var metricsPath = flag.String("metrics-path", "/metrics", "The URI path that host Prometheus metrics.")

func main() {
	ctx := context.TODO()
	flag.Parse()

	registry := prometheus.NewRegistry()
	his := services.CreateHostInfoService()
	hi, err := his.GetHostInfo(ctx)
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
			<head><title>Host Info Exporter</title></head>
			<body>
			<h1>Host Info Exporter</h1>
			<p><a href="` + *metricsPath + `">Metrics</a></p>
			</body>
			</html>`))
	})

	startupTimestamp := promauto.NewCounter(prometheus.CounterOpts{
		Name: "hostinfo_exporter_timestamp",
		ConstLabels: prometheus.Labels{
			"hostname":    hi.Hostname,
			"arch":        hi.Arch,
			"ip_vpc":      hi.VpcIp,
			"ip_external": hi.ExternalIp,
			"os":          hi.OS,
		},
	})

	startupTimestamp.Add(float64(time.Now().Unix() / 1000))

	registry.Register(startupTimestamp)

	http.Handle(*metricsPath, promhttp.HandlerFor(registry, promhttp.HandlerOpts{
		EnableOpenMetrics: true,
	}))

	fmt.Printf("Host Info Exporter running on port %s", *address)
	log.Fatal(http.ListenAndServe(*address, nil))
}
