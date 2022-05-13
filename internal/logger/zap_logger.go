package logger

import (
	"github.com/iammukeshm/structured-logging-golang-zap/utils"
	"go.uber.org/zap"
)

type zapLogger struct {
	logger *zap.Logger
}

func InitializeLogger() {
}

func createZapLogger() (Logger, error) {
	zaplogger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}
	utils.InitializeLogger()
	logger := &zapLogger{
		zaplogger,
	}

	utils.Logger.Error("Not able to reach blog.", zap.String("url", "codewithmukesh.com"))

	return logger, nil
}

func (l *zapLogger) Debug(...interface{}) error {
	utils.Logger.Info("Hello World")
	l.logger.
}

func (l *zapLogger) Info(...interface{}) error {
	utils.Logger.Info("Hello World")
	l.logger.
}

func (l *zapLogger) Error(...interface{}) error {

}

func (l *zapLogger) Fatal(...interface{}) error {
	return nil
}
