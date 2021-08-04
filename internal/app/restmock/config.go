package restmock

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Config struct {
	Interactions []Interaction `yaml:"interactions"`
}

type Interaction struct {
	Request struct {
		Path string `yaml:"path"`
		Method string `yaml:"method"`
	} `yaml:"request"`
	Response struct {
		StatusCode int    `yaml:"statusCode"`
		Body       string `yaml:"body"`
		Type       string `yaml:"type"`
	} `yaml:"response"`
}

func ParseConfig(path string) *Config{
	conf := Config{}
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("Cant read config file", err)
	}
	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		log.Fatal("Cant unmarshall config file", err)
	}
	return &conf
}