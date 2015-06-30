package nrmetrics

import (
	"github.com/neocortical/newrelic"
	metrics "github.com/rcrowley/go-metrics"
)

// AddCounterMetric Adds various metrics based on a Counter, according to the supplied
// MetricConfig
func AddCounterMetric(plugin *newrelic.Plugin, counter metrics.Counter, config MetricConfig) {
	if config.Count {
		plugin.AddMetric(newrelic.NewMetric(config.Name+"/Count", config.Unit, func() (float64, error) { return float64(counter.Snapshot().Count()), nil }))
	}
}
