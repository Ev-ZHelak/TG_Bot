package main

import (
	"TG_Bot/internal/config"
	"log"
	"time"

	tele "gopkg.in/telebot.v3"
)

func main() {
	// Загрузка конфигурации
	cfg := config.LoadTgBotConfig()

	bot, err := tele.NewBot(tele.Settings{
		Token:  cfg.Telegram.Token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
	}

	bot.Handle("/start", func(c tele.Context) error {
		return c.Send("Добро пожаловать!")
	})

	bot.Handle(tele.OnText, func(c tele.Context) error {
		return c.Send("Вы сказали: " + c.Text())
	})

	bot.Start()
}
