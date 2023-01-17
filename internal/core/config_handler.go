package core

import (
	yaml "gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	PsUtil PsUtil
}

type PsUtil struct {
	HostAll         bool `yaml:"host_all"`
	HostName        bool `yaml:"host_name"`
	Uptime          bool `yaml:"uptime"`
	Procs           bool `yaml:"procs"`
	Os              bool `yaml:"os"`
	Platform        bool `yaml:"platform"`
	PlatformFamily  bool `yaml:"platform_family"`
	PlatformVersion bool `yaml:"platform_version"`
	VirtSystem      bool `yaml:"virt_system"`
	VirtRole        bool `yaml:"virt_role"`
}

// NewConfig returns a new decoded Config struct
func NewConfig(configPath string) (*Config, error) {
	// Create config structure
	config := &Config{}

	// Open config file
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Init new YAML decode
	d := yaml.NewDecoder(file)

	// Start YAML decoding from file
	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}
