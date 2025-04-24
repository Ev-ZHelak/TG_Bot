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

	// Устанавливаем список команд
	commands := []tele.Command{
		{Text: "start", Description: "Начать работу с ботом"},
		{Text: "help", Description: "Помощь и инструкции"},
		{Text: "settings", Description: "Настройки профиля"},
	}
	err = bot.SetCommands(commands)
	if err != nil {
		log.Println("Ошибка установки команд:", err)
	}

	// Обработчики команд
	bot.Handle("/start", func(c tele.Context) error {
		return c.Send("Добро пожаловать! Введите /help для списка команд.")
	})

	bot.Handle("/help", func(c tele.Context) error {
		return c.Send("Список команд:\n/start - Начать\n/help - Помощь\n/settings - Настройки")
	})

	bot.Start()
}
