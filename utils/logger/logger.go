package logger

import "go.uber.org/zap"

type LoggerHandler interface {
	Panic(msg string)
	Error(msg string)
}

type logger struct {
	zap *zap.Logger
}

func New() (LoggerHandler, error) {
	l := &logger{}

	zap, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}
	l.zap = zap

	return l, err
}

func (l *logger) Panic(msg string) {
	l.zap.Panic(msg)
}
func (l *logger) Error(msg string) {
	l.zap.Error(msg)
}
