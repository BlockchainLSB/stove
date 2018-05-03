package config

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

// StoveConfig is a mandatory and essential configurations.
type StoveConfig struct {
	Servers []ServerConfig `yaml:"servers"`
}

// ServerConfig configures a server.
type ServerConfig struct {
	Name string `yaml:"name"`
	Port int    `yaml:"port"`
}

// LoadConfig loads configurations into a interface. Currently,
// it only supports YAML configuration.
func LoadConfig(f string, i interface{}) error {
	c, e := ioutil.ReadFile(f)

	if e != nil {
		return e
	}

	return yaml.Unmarshal(c, i)
}
