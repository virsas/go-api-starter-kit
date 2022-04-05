package logger

import (
	"fmt"
	"log"
)

type Logger struct{}

func New() (*Logger, error) {
	logger := Logger{}
	return &logger, nil
}

func (l Logger) Panic(msg string) {
	panic(msg)
}
func (l Logger) Error(msg string) {
	fmt.Println("test")
	log.Println(msg)
}
