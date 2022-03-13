package main

import (
	"log"
	"os"
	"strconv"
	"time"

	tele "gopkg.in/telebot.v3"
)

func main() {
	_, err := MakeConfig(ParsePath())
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

	b.Handle("/id", func(c tele.Context) error {
		return c.Send(strconv.FormatInt(c.Chat().ID, 10))
	})
	b.Start()
}
