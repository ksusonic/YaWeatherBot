package bot

import (
	"fmt"
	"github.com/jasonlvhit/gocron"
	"github.com/ksusonic/YaWeatherBot/config"
	"github.com/ksusonic/YaWeatherBot/internal/weather"
	"gopkg.in/telebot.v3/middleware"
	"log"
	"time"

	tele "gopkg.in/telebot.v3"
)

type Bot struct {
	Tele             *tele.Bot
	ForecastCfg      *config.Forecast
	ForecastChatId   tele.ChatID
	cronLaunch       string
	imgService       ImgService
	enableImgService bool
}

type ImgService interface {
	GetRandomImagePath() string
}

func NewBot(cfg *config.Config, imgService ImgService) *Bot {
	teleBot, err := tele.NewBot(tele.Settings{
		Token:  cfg.Token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal("Error creating bot:\n", err)
	}

	teleBot.Use(middleware.Logger())
	teleBot.Use(middleware.Whitelist(cfg.Admin, cfg.ForecastChatId))

	return &Bot{
		Tele:             teleBot,
		ForecastChatId:   tele.ChatID(cfg.ForecastChatId),
		ForecastCfg:      &cfg.ForecastConfig,
		cronLaunch:       cfg.Cron,
		imgService:       imgService,
		enableImgService: true,
	}
}

func (b *Bot) SendForecast(chatId tele.ChatID) {
	forecast, err := weather.GetForecast(b.ForecastCfg)
	if err != nil {
		fmt.Printf("Error getting forecast: %v\n", err)
		return
	}

	_, err = b.Tele.Send(chatId, forecast)
	if err != nil {
		log.Printf("Could not send forecast: %v\n", err)
		return
	}
	if b.imgService != nil {
		_, err = b.Tele.Send(chatId, &tele.Photo{
			File: tele.FromDisk(b.imgService.GetRandomImagePath()),
		})
		if err != nil {
			log.Printf("Could not send image: %v\n", err)
			return
		}
	}
}

func (b *Bot) Start() {
	b.initHandlers()
	go b.Tele.Start()

	err := gocron.Every(1).Day().At(b.cronLaunch).Do(b.SendForecast, b.ForecastChatId)
	if err != nil {
		log.Fatal(err)
	}

	_, next := gocron.NextRun()
	log.Println("Started gocron:", next)
	for <-gocron.Start() {
		_, next = gocron.NextRun()
		log.Println("Next launch of gocron:", next)
	}
}
