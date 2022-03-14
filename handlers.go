package main

import tele "gopkg.in/telebot.v3"

func FillHandlers(b *tele.Bot) {
	b.Handle("/id", GetChatId)
	b.Handle("/pwd", GetChatId)
	b.Handle("/ping", func(c tele.Context) error {
		return c.Reply("pong! 🏓")
	})
	b.Handle("/start", func(c tele.Context) error {
		if c.Sender().LanguageCode == "ru" {
			return c.Send("Привет! Я приватный бот от @ksusonic :)")
		} else {
			return c.Send("Hello! I am private bot by @ksusonic")
		}
	})
	b.Handle("/subscribe", func(c tele.Context) error {
		if c.Sender().LanguageCode == "ru" {
			return c.Send("Скоро так можно будет подписаться на обновления погоды в твоем городе!")
		} else {
			return c.Send("Soon you would be able to subscribe on weather updates!")
		}
	})
	b.Handle("/weather", func(c tele.Context) error {
		if c.Sender().LanguageCode == "ru" {
			return c.Send("Скоро так можно будет узнать погоду :)")
		} else {
			return c.Send("Soon you would be able to get weather information :)")
		}
	})
}
