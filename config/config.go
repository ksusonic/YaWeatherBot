package config

import (
	"flag"
	"fmt"
	"os"

	"github.com/caarlos0/env/v6"
	"gopkg.in/yaml.v2"
)

type Forecast struct {
	BaseUrl             string `yaml:"base_url"`
	ForecastPlaceNaming string `yaml:"forecast_place_naming"`
	Lat                 string `yaml:"lat"`
	Lon                 string `yaml:"lon"`
}

func (c Forecast) String() string {
	return fmt.Sprintf("base_url=%s, place=%s, lat=%s, lon=%s",
		c.BaseUrl, c.ForecastPlaceNaming, c.Lat, c.Lon)
}

type Config struct {
	Admin          int64  `yaml:"admin"`
	Token          string `yaml:"-" env:"TOKEN"`
	Cron           string `yaml:"cron"`
	ForecastChatId int64  `yaml:"forecast_chat_id"`
	ImgDir         string `yaml:"img_dir"`

	ForecastConfig Forecast `yaml:"forecast"`
}

func MakeConfig() (*Config, error) {
	var config Config

	path := flag.String("config", "config/config.yml", "path to bot config file")
	flag.Parse()
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
