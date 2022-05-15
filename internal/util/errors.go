package util

import (
	"fmt"

	"go.uber.org/zap"
)

func BaseError(message string, extras ...interface{}) error {
	zap.L().Sugar().Errorf(message, extras...)
	return fmt.Errorf(message, extras...)
}

func FowardingError(message string, err error) error {
	return fmt.Errorf("%w; %s", err, message)
}
