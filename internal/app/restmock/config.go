package restmock

import (
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	Interactions []Interaction `yaml:"interactions"`
}

type Interaction struct {
	Request struct {
		Path   string `yaml:"path"`
		Method string `yaml:"method"`
	} `yaml:"request"`
	Response struct {
		StatusCode int    `yaml:"statusCode"`
		Body       string `yaml:"body"`
		Type       string `yaml:"type"`
	} `yaml:"response"`
}

func ParseConfig(path string, h *HttpLogger) *Config {
	conf := Config{}
	yamlFile, err := os.ReadFile(path)
	if err != nil {
		h.logger.Fatal("Cant read config file", err)
	}
	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		h.logger.Fatal("Cant unmarshall config file", err)
	}
	return &conf
}
