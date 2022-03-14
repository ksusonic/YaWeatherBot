package main

import (
	"log"
	"os"
	"strconv"
	"time"

	tele "gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/middleware"
)

func GetChatId(c tele.Context) error {
	prefix := "Chat id: "
	id := strconv.FormatInt(c.Chat().ID, 10)
	return c.Send(prefix+id, &tele.SendOptions{
		Entities: tele.Entities{tele.MessageEntity{Type: tele.EntityCode, Offset: len(prefix), Length: len(id)}}})
}

func main() {
	config, err := MakeConfig(ParsePath())
	if err != nil {
		log.Fatal(err)
		return
	}

	pref := tele.Settings{
		Token:  os.Getenv("TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal("Check your token.\n", err)
		return
	}
	b.Use(middleware.Logger())
	adminOnly := b.Group()
	adminOnly.Use(middleware.Whitelist(config.Bot.Admin))

	FillHandlers(b)

	b.Start()
}
