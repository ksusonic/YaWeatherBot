package main

import (
	"log"

	"github.com/ksusonic/YaWeatherBot/config"
	"github.com/ksusonic/YaWeatherBot/internal/bot"
)

func main() {
	cfg, err := config.MakeConfig()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Loaded config: %s\n", cfg.ForecastConfig)

	b := bot.NewBot(cfg)

	b.Start()
}
