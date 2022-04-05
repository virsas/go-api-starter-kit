package logger

import "log"

type Logger struct{}

func InitLogger() (*Logger, error) {
	logger := Logger{}
	return &logger, nil
}

func (l Logger) Panic(msg string) {
	panic(msg)
}
func (l Logger) Error(msg string) {
	log.Println(msg)
}
