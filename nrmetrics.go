package nrmetrics

import (
	"strconv"
	"time"

	"github.com/gedex/inflector"
)

// MetricConfig defines the metrics that should be created for a given metric object,
// including the unit name and desired time precision. For example:
// Name: "MyApp/Response Time", // root of the metric name
// Unit:        "request", // the thing being timed
// Duration:    time.Millisecond, // show the histogram values in milliseconds
// Count:       true, // Send a total count metric (requests)
// Min:         true, // Send a min request time metric (milliseconds)
// Max:         true, // Send a max request time metric (milliseconds)
// Mean:        true, // Send the avg request time (milliseconds/request)
// Percentiles: []float64{0.25, 0.5, 0.75, 0.9, 0.99, 0.999}, // Send these percentiles (milliseconds/request)
// Rate1:       true, // Send the 1-minute rate, always in seconds (requests/second)
// Rate5:       true, // Send the 5-minute rate, always in seconds (requests/second)
// Rate15:      true, // Send the 15-minute rate, always in seconds (requests/second)
// RateMean:    true, // Send the mean rate, always in seconds (requests/second)
// StdDev:      true, // Send the std dev (milliseconds)
// Sum:         true, // Send the sum of all recorded times (milliseconds)
// Variance:    true, // Send the variance (milliseconds)
type MetricConfig struct {
	Name        string
	Unit        string
	Duration    time.Duration
	Count       bool
	Value       bool
	Min         bool
	Max         bool
	Mean        bool
	Percentiles []float64
	Rate1       bool
	Rate5       bool
	Rate15      bool
	RateMean    bool
	StdDev      bool
	Sum         bool
	Variance    bool
}

// parseDuration generates a human-readable string value for a Duration
func parseDuration(dur time.Duration) string {
	switch dur {
	case time.Nanosecond:
		return "nanoseconds"
	case time.Microsecond:
		return "microseconds"
	case time.Millisecond:
		return "milliseconds"
	case time.Second:
		return "seconds"
	case time.Minute:
		return "minutes"
	case time.Hour:
		return "hours"
	case time.Hour * 24:
		return "days"
	default:
		// TODO: something more intelligent here
		return strconv.FormatInt(int64(dur), 10)
	}
}

func pctToA(pct float64) string {
	pct *= 100.0
	// TODO: Ignoring ungrammatical edge cases such as 1th, 2th, 53th, etc.
	return strconv.FormatFloat(pct, 'f', -1, 64) + "th Percentile"
}

func singularize(unit string) string {
	return inflector.Singularize(unit)
}
