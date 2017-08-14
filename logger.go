package gsrlog

// Logger describes a logger-aware instance.
type Logger interface {
	// System is unusable.
	Emergency(format string, args ...interface{})
	// Action must be taken immediately.
	Alert(format string, args ...interface{})
	// Critical conditions.
	Critical(format string, args ...interface{})
	// Runtime errors that do not require immediate action but should typically be logged and monitored.
	Error(format string, args ...interface{})
	// Exceptional occurrences that are not errors.
	Warning(format string, args ...interface{})
	// Normal but significant events.
	Notice(format string, args ...interface{})
	// Interesting events.
	Info(format string, args ...interface{})
	// Detailed debug information.
	Debug(format string, args ...interface{})
	// Logs with an arbitrary level.
	Log(level Level, format string, args ...interface{})
}
