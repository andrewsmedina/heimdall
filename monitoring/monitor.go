package monitoring

// Monitor is the basic interface of this package.
type Monitor interface {
	// AddAlarm adds a new alarm.
	AddAlarm() error

	// DeleteAlarm deletes an alarm.
	DeleteAlarm() error

	// ListAlarms lists alarms.
	ListAlarms()

	// ListMetrics list metrics.
	ListMetrics()
}
