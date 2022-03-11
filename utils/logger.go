package utils

import (
	"go.uber.org/zap"
)

func InitLogger() (*zap.Logger, error) {
	var logger *zap.Logger

	logger, err := zap.NewProduction()
	logger.Info("Logger initialized...")

	return logger, err
}
