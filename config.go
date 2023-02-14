package main

import (
	"flag"
	"os"

	"github.com/caarlos0/env/v6"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Admin int64  `yaml:"admin"`
	Token string `yaml:"-" env:"TOKEN"`
}

func MakeConfig(configPath string) (*Config, error) {
	var config Config

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)
	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	if err := env.Parse(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

func ParsePath() string {
	var path string
	flag.StringVar(&path, "config", "./config.yml", "path to bot config file")
	flag.Parse()
	return path
}
