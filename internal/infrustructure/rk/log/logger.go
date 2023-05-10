package log

import (
	"go.uber.org/zap"
)

var sl *zap.SugaredLogger

func New(s *zap.SugaredLogger) {
	sl = s
}

// Debug uses fmt.Sprint to construct and log a message.
func Debug(args ...interface{}) {
	sl.Debug(args)
}

// Info uses fmt.Sprint to construct and log a message.
func Info(args ...interface{}) {
	sl.Info(args)
}

// Warn uses fmt.Sprint to construct and log a message.
func Warn(args ...interface{}) {
	sl.Warn(args)
}

// Error uses fmt.Sprint to construct and log a message.
func Error(args ...interface{}) {
	sl.Error(args)
}

// Errorf uses fmt.Sprintf to log a templated message.
func Errorf(template string, args ...interface{}) {
	sl.Errorf(template, args)
}
