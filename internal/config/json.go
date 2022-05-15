package config

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/renatobrittoaraujo/rendering/internal/util"
	"go.uber.org/zap"
)

type jsonConfig struct {
	data   map[string]interface{}
	logger *zap.Logger
}

func (c *jsonConfig) load(configSource *ConfigSource, logger *zap.Logger) error {
	jsonFile, err := os.Open(configSource.Address)
	if err != nil {
		return util.FowardingError("failed to open json config file", err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal([]byte(byteValue), &c.data)

	c.logger = logger
	return nil
}

func (c *jsonConfig) Get(target string) (string, error) {
	if val, ok := c.data[target]; ok {
		if sval, ok := val.(string); ok {
			return sval, nil
		}

		return "", util.BaseError(c.logger, "failed to convert config var to string '%s'", target)
	}

	return "", util.BaseError(c.logger, "failed to find config var '%s'", target)
}
