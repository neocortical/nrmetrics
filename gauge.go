package nrmetrics

import (
	"github.com/neocortical/newrelic"
	metrics "github.com/rcrowley/go-metrics"
)

// AddGaugeMetric Adds various metrics based on a Gauge, according to the supplied
// MetricConfig.
func AddGaugeMetric(plugin *newrelic.Plugin, gauge metrics.Gauge, config MetricConfig) {
	if config.Count {
		plugin.AddMetric(newrelic.NewMetric(config.Name+"/Value", config.Unit, func() (float64, error) { return float64(gauge.Snapshot().Value()), nil }))
	}
}
