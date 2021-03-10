package logger

import (
	"log"
	"os"
)

var Log = newLogger()

type Logger interface {
	Info(msg string)
	Error(msg string)
}

type logger struct {
	iLog *log.Logger
	eLog *log.Logger
}

func newLogger() Logger {
	return &logger{
		iLog: log.New(os.Stdout, "[INFO] ", log.Lmicroseconds),
		eLog: log.New(os.Stdout, "[FATE] ", log.Lmicroseconds),
	}
}

func (l *logger) Info(msg string) {
	l.iLog.Println(msg)
}

func (l *logger) Error(msg string) {
	l.eLog.Println(msg)
}
