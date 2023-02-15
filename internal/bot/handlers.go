package bot

import (
	"fmt"
	"github.com/ksusonic/YaWeatherBot/config"
	"github.com/ksusonic/YaWeatherBot/internal/weather"
	"strconv"

	tele "gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/middleware"
)

func (b *Bot) initHandlers() {
	b.Tele.Handle("/pwd", func(c tele.Context) error {
		prefix := "Chat id: "
		id := strconv.FormatInt(c.Chat().ID, 10)
		return c.Send(prefix+id, &tele.SendOptions{
			Entities: tele.Entities{tele.MessageEntity{Type: tele.EntityCode, Offset: len(prefix), Length: len(id)}}})
	})
	b.Tele.Handle("/ping", func(c tele.Context) error {
		return c.Reply("pong! 🏓")
	})
	b.Tele.Handle("/weather", func(c tele.Context) error {
		forecast, err := weather.GetForecast(b.ForecastCfg)
		if err != nil {
			fmt.Printf("Error getting forecast: %v\n", err)
			return err
		}
		return c.Reply(forecast)
	})
}

func initMiddleware(b *tele.Bot, config *config.Config) {
	b.Use(middleware.Logger())

	if config.Admin != 0 {
		adminOnly := b.Group()
		adminOnly.Use(middleware.Whitelist(config.Admin))
	}
}
