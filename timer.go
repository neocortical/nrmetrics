package nrmetrics

import (
	"github.com/neocortical/newrelic"
	metrics "github.com/rcrowley/go-metrics"
)

// AddTimerMetric Adds various metrics based on a Timer, according to the supplied
// MetricConfig
func AddTimerMetric(plugin *newrelic.Plugin, timer metrics.Timer, config MetricConfig) {
	durString := parseDuration(config.Duration)
	unitSingular := singularize(config.Unit)
	units := durString + "/" + unitSingular
	if config.Count {
		plugin.AddMetric(newrelic.NewMetric(config.Name+"/Totals/Count", config.Unit, func() (float64, error) { return float64(timer.Snapshot().Count()), nil }))
	}
	if config.Min {
		plugin.AddMetric(newrelic.NewMetric(config.Name+"/Stats/Min", durString, func() (float64, error) { return float64(timer.Snapshot().Min() / int64(config.Duration)), nil }))
	}
	if config.Max {
		plugin.AddMetric(newrelic.NewMetric(config.Name+"/Stats/Max", durString, func() (float64, error) { return float64(timer.Snapshot().Max() / int64(config.Duration)), nil }))
	}
	if config.Mean {
		plugin.AddMetric(newrelic.NewMetric(config.Name+"/Time/Mean", units, func() (float64, error) { return timer.Snapshot().Mean() / float64(config.Duration), nil }))
	}
	for _, pct := range config.Percentiles {
		if pct == 0.5 {
			localPct := pct
			plugin.AddMetric(newrelic.NewMetric(config.Name+"/Time/Median", units, func() (float64, error) { return timer.Snapshot().Percentile(localPct) / float64(config.Duration), nil }))
		} else {
			localPct := pct
			plugin.AddMetric(newrelic.NewMetric(config.Name+"/Time/"+pctToA(pct), units, func() (float64, error) { return timer.Snapshot().Percentile(localPct) / float64(config.Duration), nil }))
		}
	}
	if config.Rate1 {
		plugin.AddMetric(newrelic.NewMetric(config.Name+"/Rate/1 Min. Rate", config.Unit+"/second", func() (float64, error) { return timer.Snapshot().Rate1(), nil }))
	}
	if config.Rate5 {
		plugin.AddMetric(newrelic.NewMetric(config.Name+"/Rate/5 Min. Rate", config.Unit+"/second", func() (float64, error) { return timer.Snapshot().Rate5(), nil }))
	}
	if config.Rate15 {
		plugin.AddMetric(newrelic.NewMetric(config.Name+"/Rate/15 Min. Rate", config.Unit+"/second", func() (float64, error) { return timer.Snapshot().Rate15(), nil }))
	}
	if config.RateMean {
		plugin.AddMetric(newrelic.NewMetric(config.Name+"/Rate/Mean Rate", config.Unit+"/second", func() (float64, error) { return timer.Snapshot().RateMean(), nil }))
	}
	if config.StdDev {
		plugin.AddMetric(newrelic.NewMetric(config.Name+"/Stats/Std Dev", durString, func() (float64, error) { return timer.Snapshot().StdDev() / float64(config.Duration), nil }))
	}
	if config.Sum {
		plugin.AddMetric(newrelic.NewMetric(config.Name+"/Totals/Sum", durString, func() (float64, error) { return float64(timer.Snapshot().Sum() / int64(config.Duration)), nil }))
	}
	if config.Variance {
		plugin.AddMetric(newrelic.NewMetric(config.Name+"/Stats/Variance", durString, func() (float64, error) { return timer.Snapshot().Variance() / float64(config.Duration), nil }))
	}
}
