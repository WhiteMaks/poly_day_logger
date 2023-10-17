package logger

type PolyDayLogger struct {
	logger Logger
	name   string
}

func NewPolyDayLogger(logger Logger, name string) *PolyDayLogger {
	return &PolyDayLogger{
		logger: logger,
		name:   name,
	}
}

func (logger *PolyDayLogger) Error(message Message) {
	logger.fillMessage(&message, 3, ERROR.String())
	logger.logger.Error(message)
}

func (logger *PolyDayLogger) Warn(message Message) {
	logger.fillMessage(&message, 4, WARNING.String())
	logger.logger.Warn(message)
}

func (logger *PolyDayLogger) Info(message Message) {
	logger.fillMessage(&message, 6, INFO.String())
	logger.logger.Info(message)
}

func (logger *PolyDayLogger) Debug(message Message) {
	logger.fillMessage(&message, 7, DEBUG.String())
	logger.logger.Debug(message)
}

func (logger *PolyDayLogger) fillMessage(message *Message, level int32, levelName string) {
	message.Level = level
	message.LevelName = levelName
	message.Version = "1.1"
	message.Host = logger.name
	message.Facility = logger.name
}
