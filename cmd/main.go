package main

import (
	"github.com/ksusonic/YaWeatherBot/internal/img"
	"log"

	"github.com/ksusonic/YaWeatherBot/config"
	"github.com/ksusonic/YaWeatherBot/internal/bot"
)

func main() {
	cfg, err := config.MakeConfig()
	if err != nil {
		log.Fatal(err)
	}
	imgService, err := img.NewImg(cfg.ImgDir)
	if err != nil {
		log.Println("ImgService will not be used:", err)
	}

	log.Printf("Loaded config: %s\n", cfg.ForecastConfig)

	b := bot.NewBot(cfg, imgService)

	b.Start()
}
