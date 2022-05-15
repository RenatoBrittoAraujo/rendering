package logger

import (
	"fmt"

	"go.uber.org/zap"
)

// test
func CreateLogger() (*zap.Logger, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, fmt.Errorf("%w; failed to create zap logger", err)
	}

	defer logger.Sync()
	return logger, nil
}
