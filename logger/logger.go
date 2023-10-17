package logger

import (
	"fmt"
	"log"
)

type Logger interface {
	GetLevel() Level
	Error(message any) bool
	Warn(message any) bool
	Info(message any) bool
	Debug(message any) bool
}

type Default struct {
	level Level
}

func NewDefaultLogger(level Level) *Default {
	return &Default{
		level: level,
	}
}

func (logger *Default) Error(message any) bool {
	if logger.level >= ERROR {
		logger.Println(message, ERROR)
		return true
	}

	return false
}

func (logger *Default) Warn(message any) bool {
	if logger.level >= WARNING {
		logger.Println(message, WARNING)
		return true
	}

	return false
}

func (logger *Default) Info(message any) bool {
	if logger.level >= INFO {
		logger.Println(message, INFO)
		return true
	}

	return false
}

func (logger *Default) Debug(message any) bool {
	if logger.level >= DEBUG {
		logger.Println(message, DEBUG)
		return true
	}

	return false
}

func (logger *Default) GetLevel() Level {
	return logger.level
}

func (logger *Default) Println(message any, level Level) {
	log.Println("[ " + level.String() + " ] " + fmt.Sprintf("%v", message))
}
