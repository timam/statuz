package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/timam/statuz/config"
	"net/http"
)

func SetWebpagePrometheusMetric(url string, status int64) {
	metric := config.MetricName()
	gauge := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: metric,
			Help: "Indicates whether the webpage is up (1) or down (0)",
		},
		[]string{"url"}, // Add a "url" label
	)

	if status == http.StatusOK {
		gauge.WithLabelValues(url).Set(float64(1)) // Up
	} else {
		gauge.WithLabelValues(url).Set(float64(0)) // Down
	}

	// Register the gauge metric with Prometheus.
	prometheus.Register(gauge)
}
