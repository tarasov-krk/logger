package logger

import (
	"log"
	"os"
)

type StdLogger struct {
	logger *log.Logger
	level  Level
}

func NewStd() *StdLogger {
	return &StdLogger{
		logger: log.New(os.Stderr, "", log.LstdFlags),
		level:  DefaultLogLevel,
	}
}

func (l *StdLogger) SetLevel(level Level) {
	l.level = level
}

func (l *StdLogger) GetLevel() Level {
	return l.level
}

func (l *StdLogger) Write(level string, msg ...any) {
	l.logger.Println(level, msg)
}
