package logger

type Level int32

const (
	ERROR   Level = 0
	WARNING Level = 1
	INFO    Level = 2
	DEBUG   Level = 3
)

func (l Level) String() string {
	switch l {
	case ERROR:
		return "ERROR"
	case WARNING:
		return "WARNING"
	case INFO:
		return "INFO"
	case DEBUG:
		return "DEBUG"
	}
	return "unknown"
}
