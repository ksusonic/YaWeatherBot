package main

import (
	"log"

	"github.com/ksusonic/YaWeatherBot/internal/bot"
)

func main() {
	config, err := bot.MakeConfig()
	if err != nil {
		log.Fatal(err)
	}

	b := bot.NewBot(config)

	b.Start()
}
