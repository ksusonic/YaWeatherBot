package main

import (
	"flag"
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	Bot struct {
		Admin string `yaml:"admin"`
	}
}

func MakeConfig(configPath string) (*Config, error) {
	config := &Config{}

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)

	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}

func ParsePath() string {
	var path string
	flag.StringVar(&path, "config", "./config.yml", "path to bot config file")
	flag.Parse()
	return path
}
