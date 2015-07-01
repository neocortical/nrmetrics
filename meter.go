package nrmetrics

import (
	"github.com/neocortical/newrelic"
	metrics "github.com/rcrowley/go-metrics"
)

// AddMeterMetric Adds various metrics based on a Meter, according to the supplied
// MetricConfig
func AddMeterMetric(plugin *newrelic.Plugin, meter metrics.Meter, config MetricConfig) {
	if config.Count {
		plugin.AddMetric(newrelic.NewMetric(config.Name+"/Totals/Count", config.Unit, func() (float64, error) { return float64(meter.Snapshot().Count()), nil }))
	}
	if config.Rate1 {
		plugin.AddMetric(newrelic.NewMetric(config.Name+"/Rate/1 Min. Rate", config.Unit+"/second", func() (float64, error) { return meter.Snapshot().Rate1(), nil }))
	}
	if config.Rate5 {
		plugin.AddMetric(newrelic.NewMetric(config.Name+"/Rate/5 Min. Rate", config.Unit+"/second", func() (float64, error) { return meter.Snapshot().Rate5(), nil }))
	}
	if config.Rate15 {
		plugin.AddMetric(newrelic.NewMetric(config.Name+"/Rate/15 Min. Rate", config.Unit+"/second", func() (float64, error) { return meter.Snapshot().Rate15(), nil }))
	}
	if config.RateMean {
		plugin.AddMetric(newrelic.NewMetric(config.Name+"/Rate/Mean Rate", config.Unit+"/second", func() (float64, error) { return meter.Snapshot().RateMean(), nil }))
	}
}
