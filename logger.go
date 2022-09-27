package logger

import (
	"fmt"
	"runtime"
)

type Level int

const (
	LevelError Level = iota
	LevelWarning
	LevelInfo
	LevelDebug
)

var DefaultLogLevel = LevelError

func (l Level) IsValid() bool {
	switch l {
	case LevelError, LevelWarning, LevelInfo, LevelDebug:
		return true
	default:
		return false
	}
}

type Logger interface {
	Write(level string, msg ...any)
	SetLevel(l Level)
	GetLevel() Level
}

// Default set StdOut logger
var logger Logger = NewStd()

func SetLevel(level Level) {
	if !level.IsValid() {
		return
	}
	logger.SetLevel(level)
}

func SetLogger(l Logger) {
	logger = l
}

func Info(msg ...any) {
	if logger.GetLevel() <= LevelInfo {
		logger.Write("INFO", appendContext(msg...)...)
	}
}

func Warn(msg ...any) {
	if logger.GetLevel() <= LevelWarning {
		logger.Write("WARNING", appendContext(msg...)...)
	}
}

func Error(msg ...any) {
	if logger.GetLevel() <= LevelError {
		logger.Write("ERROR", appendContext(msg...)...)
	}
}

func Debug(msg ...any) {
	if logger.GetLevel() <= LevelDebug {
		logger.Write("DEBUG", appendContext(msg...)...)
	}
}

func currentFunction() string {
	counter, _, line, _ := runtime.Caller(3)
	return runtime.FuncForPC(counter).Name() + fmt.Sprint(": ", line)
}

func appendContext(msg ...any) []any {
	return append(msg, " ["+currentFunction()+"]")
}
