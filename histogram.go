package nrmetrics

import (
	"github.com/neocortical/newrelic"
	metrics "github.com/rcrowley/go-metrics"
)

// AddHistogramMetric Adds various metrics based on a Histogram, according to the supplied
// MetricConfig
func AddHistogramMetric(plugin *newrelic.Plugin, histo metrics.Histogram, config MetricConfig) {
	durString := parseDuration(config.Duration)
	unitSingular := singularize(config.Unit)
	units := durString + "/" + unitSingular
	if config.Count {
		plugin.AddMetric(newrelic.NewMetric(config.Name+"/Totals/Count", config.Unit, func() (float64, error) { return float64(histo.Snapshot().Count()), nil }))
	}
	if config.Min {
		plugin.AddMetric(newrelic.NewMetric(config.Name+"/Stats/Min", durString, func() (float64, error) { return float64(histo.Snapshot().Min() / int64(config.Duration)), nil }))
	}
	if config.Max {
		plugin.AddMetric(newrelic.NewMetric(config.Name+"/Stats/Max", durString, func() (float64, error) { return float64(histo.Snapshot().Max() / int64(config.Duration)), nil }))
	}
	if config.Mean {
		plugin.AddMetric(newrelic.NewMetric(config.Name+"/Rate/Mean", units, func() (float64, error) { return histo.Snapshot().Mean() / float64(config.Duration), nil }))
	}
	for _, pct := range config.Percentiles {
		if pct == 0.5 {
			localPct := pct
			plugin.AddMetric(newrelic.NewMetric(config.Name+"/Rate/Median", units, func() (float64, error) { return histo.Snapshot().Percentile(localPct) / float64(config.Duration), nil }))
		} else {
			localPct := pct
			plugin.AddMetric(newrelic.NewMetric(config.Name+"/Rate/"+pctToA(pct), units, func() (float64, error) { return histo.Snapshot().Percentile(localPct) / float64(config.Duration), nil }))
		}
	}
	if config.StdDev {
		plugin.AddMetric(newrelic.NewMetric(config.Name+"/Stats/Std Dev", durString, func() (float64, error) { return histo.Snapshot().StdDev() / float64(config.Duration), nil }))
	}
	if config.Sum {
		plugin.AddMetric(newrelic.NewMetric(config.Name+"/Totals/Sum", durString, func() (float64, error) { return float64(histo.Snapshot().Sum() / int64(config.Duration)), nil }))
	}
	if config.Variance {
		plugin.AddMetric(newrelic.NewMetric(config.Name+"/Stats/Variance", durString, func() (float64, error) { return histo.Snapshot().Variance() / float64(config.Duration), nil }))
	}
}
