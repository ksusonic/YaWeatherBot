package bot

import (
	"log"
	"time"

	tele "gopkg.in/telebot.v3"
)

type Bot struct {
	Tele *tele.Bot
}

func NewBot(config *Config) *Bot {
	teleBot, err := tele.NewBot(tele.Settings{
		Token:  config.Token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal("Error creating bot:\n", err)
	}

	initHandlers(teleBot)
	initMiddleware(teleBot, config)

	return &Bot{
		Tele: teleBot,
	}
}

func (b Bot) Start() {
	b.Tele.Start()
}
