package main

import (
	"net/http"
	"os"

	"github.com/aburtasov/jellyfin_exporter/exporter"
	kingpin "github.com/alecthomas/kingpin/v2"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"log"

	"github.com/prometheus/common/version"
)

var (
	apiUrl        = kingpin.Flag("jellyfin.apiurl", "Jellyfin API url.").Default(getEnv("JELLYFIN_APIURL", "http://127.0.0.1")).String()
	apiKey        = kingpin.Flag("jellyfin.apikey", "ApiKey for jellyfin API.").Default(getEnv("JELLYFIN_APIKEY", "")).String()
	timeout       = kingpin.Flag("jellyfin.timeout", "jellyfin connect timeout.").Default("1s").Duration()
	listenAddress = kingpin.Flag("web.listen-address", "Address to listen on for web interface and telemetry.").Default(getEnv("WEB_LISTEN_ADDRESS", ":9249")).String()
	metricsPath   = kingpin.Flag("web.telemetry-path", "Path under which to expose metrics.").Default("/metrics").String()
)

func main() {

	log.Flags()
	kingpin.Version(version.Print("jellyfin_exporter"))
	kingpin.HelpFlag.Short('h')
	kingpin.Parse()

	log.Println("Starting jellyfin_exporter", version.Info())
	log.Println("Build context", version.BuildContext())

	if *apiKey == "" {
		log.Fatal("Empty api key")
	}

	prometheus.MustRegister(exporter.New(*apiUrl, *apiKey, *timeout))

	http.Handle(*metricsPath, promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(`<html>
      <head><title>Jellyfin Exporter</title></head>
      <body>
      <h1>Jellyfin Exporter</h1>
      <p><a href='` + *metricsPath + `'>Metrics</a></p>
      </body>
      </html>`))
		if err != nil {

			log.Fatal(err)
		}
	})
	log.Println("Starting HTTP server on", *listenAddress)
	log.Fatal(http.ListenAndServe(*listenAddress, nil))

}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
