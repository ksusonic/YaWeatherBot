package bot

import (
	"fmt"
	"github.com/jasonlvhit/gocron"
	"github.com/ksusonic/YaWeatherBot/config"
	"github.com/ksusonic/YaWeatherBot/internal/weather"
	"log"
	"time"

	tele "gopkg.in/telebot.v3"
)

type Bot struct {
	Tele           *tele.Bot
	ForecastCfg    *config.Forecast
	ForecastChatId tele.ChatID
	cronLaunch     string
}

func NewBot(cfg *config.Config) *Bot {
	teleBot, err := tele.NewBot(tele.Settings{
		Token:  cfg.Token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal("Error creating bot:\n", err)
	}

	initMiddleware(teleBot, cfg)

	return &Bot{
		Tele:           teleBot,
		ForecastChatId: tele.ChatID(cfg.ForecastChatId),
		ForecastCfg:    &cfg.ForecastConfig,
		cronLaunch:     cfg.Cron,
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
