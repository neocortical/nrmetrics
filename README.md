# nrmetrics - Apply rcrowley/go-metrics to NewRelic plugins

This package complements neocortical/newrelic

## Install

`go get github.com/neocortical/nrmetrics`

## Usage (with the `newrelic` package)

```go
import (
	"time"

	"github.com/neocortical/newrelic"
	nrmetrics "github.com/neocortical/newrelic_metrics"
	metrics "github.com/rcrowley/go-metrics"
)

nr := newrelic.New("abc123")

timer := metrics.NewTimer()

plugin := &newrelic.Plugin{
	Name: "My Awesome Plugin",
	GUID: "net.neocortical.foo.bar.MyAwesomePlugin",
}

// All timer metrics turned on in this example. Only set what you want to true
metricConf := nrmetrics.MetricConfig{
	Name:        "App/Response Time",
	Unit:        "request",
	Duration:    time.Millisecond,
	Count:       true,
	Min:         true,
	Max:         true,
	Mean:        true,
	Percentiles: []float64{0.25, 0.5, 0.75, 0.9, 0.99, 0.999},
	Rate1:       true,
	Rate5:       true,
	Rate15:      true,
	RateMean:    true,
	StdDev:      true,
	Sum:         true,
	Variance:    true,
}

// Add your timer to your NewRelic plugin
nrmetrics.AddTimerMetric(plugin, timer, metricConf)
nr.AddPlugin(plugin)

// Start sending metrics data to NewRelic
nr.Run()
```
