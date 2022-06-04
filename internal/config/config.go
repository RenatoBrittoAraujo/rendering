package config

import (
	"github.com/renatobrittoaraujo/rendering/internal/util"
	"go.uber.org/zap"
)

type ConfigSourceType string

const (
	JSON ConfigSourceType = "json"
)

type Config interface {
	load(configSource *ConfigSource, logger *zap.Logger) error
	Get(target string) (string, error)
}

type ConfigSource struct {
	SourceType     ConfigSourceType
	Address        string
	Authentication string
}

func LoadConfigFromSource(source *ConfigSource, logger *zap.Logger) (Config, error) {
	var config Config

	switch source.SourceType {
	case "json":
		config = &jsonConfig{}
	default:
		return nil, util.BaseError(logger, "failed to locate config source '%s'", string(source.SourceType))
	}

	err := config.load(source, logger)
	if err != nil {
		return nil, err
	}
	return config, nil
}
