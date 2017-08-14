package gsrlog

// LoggerAware describes a logger-aware struct
type LoggerAware interface {
	// SetLogger sets a logger
	SetLogger(logger Logger)
}
