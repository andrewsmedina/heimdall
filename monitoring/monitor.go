package monitoring

// Alarm represents an alarm.
type Alarm struct{}

// Metric represents a metric.
type Metric struct{}

// Monitor is the basic interface of this package.
type Monitor interface {
	// AddAlarm adds a new alarm.
	AddAlarm() error

	// DeleteAlarm deletes an alarm.
	DeleteAlarm() error

	// ListAlarms lists alarms.
	ListAlarms() []Alarm

	// ListMetrics list metrics.
	ListMetrics() []Metric
}
