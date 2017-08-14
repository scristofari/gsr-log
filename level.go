package gsrlog

// Level specify the type.
type Level string

// Describes log levels
const (
	DEBUG     Level = "debug"
	INFO      Level = "info"
	NOTICE    Level = "notice"
	WARNING   Level = "warning"
	ERROR     Level = "error"
	CRITICAL  Level = "critical"
	ALERT     Level = "alert"
	EMERGENCY Level = "emergency"
)
