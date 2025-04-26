package handlers

import (
	"fmt"

	tele "gopkg.in/telebot.v3"
)

// Обработчики команд
func MainHandlerCommands(b *tele.Bot) {
	b.Handle("/start", func(c tele.Context) error {
		return c.Send("Добро пожаловать! Введите /help для списка команд.")
	})

	b.Handle("/help", func(c tele.Context) error {
		fmt.Println(c.Message())
		return c.Send("Список команд:\n/start - Начать\n/help - Помощь\n/settings - Настройки")
	})
}
