package util

import (
	"fmt"

	"go.uber.org/zap"
)

func BaseError(logger *zap.Logger, message string, extras ...interface{}) error {
	logger.Sugar().Errorf(message, extras...)
	return fmt.Errorf(message, extras...)
}

func FowardingError(message string, err error) error {
	return fmt.Errorf("%s; %w", message, err)
}
