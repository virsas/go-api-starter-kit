package logger

import (
	"go.uber.org/zap"
)

var logger *zap.Logger

func initZap() error {
	var err error

	logger, err = zap.NewProduction()
	if err != nil {
		return err
	}

	return nil
}
