package logger

import (
	"encoding/json"
	"net"
)

type Gelf struct {
	logger Logger
	conn   net.Conn
}

type Message struct {
	Version      string `json:"version"`       //GELF spec version – “1.1”
	Host         string `json:"host"`          //the name of the host, source or application that sent this message
	ShortMessage string `json:"short_message"` //a short descriptive message
	FullMessage  string `json:"full_message"`  //a long message that can i.e. contain a backtrace
	Timestamp    int64  `json:"timestamp"`     //seconds since UNIX epoch with optional decimal places for milliseconds
	Level        int32  `json:"level"`         //the level equal to the standard syslog levels
	Facility     string `json:"facility"`      //optional, deprecated. Send as additional field instead
	UserId       int64  `json:"__user_id"`     //user id
	WsId         string `json:"__ws_id"`       //web socket id
	LevelName    string `json:"__level_name"`  //level name
}

func NewGelf(logger Logger, address string) *Gelf {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		panic("GELF not ready: " + err.Error())
	}

	return &Gelf{
		logger: logger,
		conn:   conn,
	}
}

func (logger *Gelf) Error(message any) bool {
	messageString := logger.getMessage(message)

	isPrinted := logger.logger.Error(messageString)
	if isPrinted {
		logger.sendMessage(messageString)
	}

	return isPrinted
}

func (logger *Gelf) Warn(message any) bool {
	messageString := logger.getMessage(message)

	isPrinted := logger.logger.Warn(messageString)
	if isPrinted {
		logger.sendMessage(messageString)
	}

	return isPrinted
}

func (logger *Gelf) Info(message any) bool {
	messageString := logger.getMessage(message)

	isPrinted := logger.logger.Info(messageString)
	if isPrinted {
		logger.sendMessage(messageString)
	}

	return isPrinted
}

func (logger *Gelf) Debug(message any) bool {
	messageString := logger.getMessage(message)

	isPrinted := logger.logger.Debug(messageString)
	if isPrinted {
		logger.sendMessage(messageString)
	}

	return isPrinted
}

func (logger *Gelf) GetLevel() Level {
	return logger.logger.GetLevel()
}

func (logger *Gelf) sendMessage(message string) {
	logger.conn.Write([]byte(message + "\n"))
}

func (logger *Gelf) getMessage(message any) string {
	bytes, _ := json.Marshal(message)
	return string(bytes)
}
