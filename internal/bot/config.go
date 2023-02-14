package bot

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

func MakeConfig() (*Config, error) {
	var config Config

	path := flag.String("config", "config/config.yml", "path to bot config file")

	file, err := os.ReadFile(*path)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}

	if err := env.Parse(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
