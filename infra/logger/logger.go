package logger

import (
	"log"
	"os"

	"github.com/shinnosuke-K/Tech-Train-CA-Go/domain/repository"
)

var Log = newLogger()

type logger struct {
	iLog *log.Logger
	eLog *log.Logger
}

func newLogger() repository.Logger {
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
