package handlers

import (
	tele "gopkg.in/telebot.v3"
)

// Обработчики команд
func MainHandlerCommands(b *tele.Bot) {
	b.Handle("/start", func(c tele.Context) error {
		return c.Send(`🏆 Welcome to Grey Matter Games!!! 🏆

🤖💬 Здесь сходятся самые азартные умники желающие сразиться в честном бою?

Ты сможешь:
💥 Проверить свою эрудицию
💥 Прокачать знания
💥 Возможность бить рекорды
💥 Весело провести время!

🤖📚 Моя база игр постоянно растёт, так что скучно не будет.😊

➡ Выбери игру и начнём? 👉 /games`)
	})

	b.Handle("/help", func(c tele.Context) error {
		return c.Send("Список команд:\n/start - Начать\n/help - Помощь\n/settings - Настройки")
	})
}
