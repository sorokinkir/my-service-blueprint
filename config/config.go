package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

// Config struct for api rest service
type Config struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

// NewConfig returns a new decoded Config struct
func NewConfig(configPath string) (*Config, error) {
	// create config structure
	config := &Config{}

	// open config file
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Init new yaml decode
	d := yaml.NewDecoder(file)

	// Start YAML decoding from file
	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}
