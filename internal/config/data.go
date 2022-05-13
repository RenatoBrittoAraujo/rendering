package config

type ConfigSourceType int

const (
	JSON ConfigSourceType = iota
)

type Config interface {
	load() error
	Get(target string) string
}

type ConfigSource struct {
	SourceType     ConfigSourceType
	Address        string
	Authentication string
}

func LoadConfigFromSource(source *ConfigSource) (Config, error) {
	var config Config

	switch source.SourceType {
	case JSON:
		config = &jsonConfig{}
	default:
		return nil, 
	}

	err := config.load()
	if err != nil {
		return nil, err
	}
	return config, nil
}

// Implementations

type jsonConfig struct {
}

func (c *jsonConfig) load() error {
	return nil
}

func (c *jsonConfig) Get(target string) string {
	return target
}
