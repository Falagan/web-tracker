package pkg

type LogLevel struct {
	name string
}

var (
	LogLevelInfo  = LogLevel{name: "INFO"}
	LogLevelWarn  = LogLevel{name: "WARNING"}
	LogLevelError = LogLevel{name: "ERROr"}
)

type Observer interface {
	Log(level LogLevel, msg string)
}
