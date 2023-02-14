package bot

import (
	"github.com/ksusonic/YaWeatherBot/config"
	"strconv"

	tele "gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/middleware"
)

func initHandlers(b *tele.Bot) {
	b.Handle("/pwd", func(c tele.Context) error {
		prefix := "Chat id: "
		id := strconv.FormatInt(c.Chat().ID, 10)
		return c.Send(prefix+id, &tele.SendOptions{
			Entities: tele.Entities{tele.MessageEntity{Type: tele.EntityCode, Offset: len(prefix), Length: len(id)}}})
	})
	b.Handle("/ping", func(c tele.Context) error {
		return c.Reply("pong! 🏓")
	})
}

func initMiddleware(b *tele.Bot, config *config.Config) {
	b.Use(middleware.Logger())

	if config.Admin != 0 {
		adminOnly := b.Group()
		adminOnly.Use(middleware.Whitelist(config.Admin))
	}
}
