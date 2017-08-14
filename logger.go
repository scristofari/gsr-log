package gsrlog

// Logger describes a logger-aware instance.
type Logger interface {
	// System is unusable.
	Emergency(message string, context ...interface{})
	// Action must be taken immediately.
	Alert(message string, context ...interface{})
	// Critical conditions.
	Critical(message string, context ...interface{})
	// Runtime errors that do not require immediate action but should typically be logged and monitored.
	Error(message string, context ...interface{})
	// Exceptional occurrences that are not errors.
	Warning(message string, context ...interface{})
	// Normal but significant events.
	Notice(message string, context ...interface{})
	// Interesting events.
	Info(message string, context ...interface{})
	// Detailed debug information.
	Debug(message string, context ...interface{})
	// Logs with an arbitrary level.
	Log(level Level, message string, context ...interface{})
}
