package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"

	"github.com/sudoblockio/icon-transformer/config"
)

func updateLabels(labels map[string]string) map[string]string {
	var labels_ = map[string]string{
		"network_name": config.Config.NetworkName,
	}
	for m, n := range labels {
		labels_[m] = n
	}
	return labels_
}

func CreateGuage(name string, help string, labels map[string]string) prometheus.Gauge {
	var metric = promauto.NewGauge(prometheus.GaugeOpts{
		Name:        name,
		Help:        help,
		ConstLabels: updateLabels(labels),
	})
	return metric
}

func CreateCounter(name string, help string, labels map[string]string) prometheus.Counter {
	var metric = promauto.NewCounter(prometheus.CounterOpts{
		Name:        name,
		Help:        help,
		ConstLabels: updateLabels(labels),
	})
	return metric
}

func Start() {
	// Start server
	http.Handle(config.Config.MetricsPrefix, promhttp.Handler())
	go http.ListenAndServe(":"+config.Config.MetricsPort, nil)
	zap.S().Info("Started Metrics:", config.Config.MetricsPort)
}
