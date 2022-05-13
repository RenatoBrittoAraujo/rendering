package logger

import (
	"github.com/iammukeshm/structured-logging-golang-zap/utils"
	"go.uber.org/zap"
)

func CreateLogger() (*zap.Logger, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	utils.InitializeLogger()
	utils.Logger.Info("Logger has been initialized.")
	return logger, nil
}
