package menu

import (
	"log"

	tele "gopkg.in/telebot.v3"
)

var (
	commands = []tele.Command{
		{Text: "start", Description: "Запустить бот"},
		{Text: "games", Description: "Список игр"},
		{Text: "settings", Description: "Настройки профиля"},
		{Text: "help", Description: "Помощь и инструкции"},
	}
)

// Устанавливаем список команд
func CreateMenu(b *tele.Bot) {
	err := b.SetCommands(commands)
	if err != nil {
		log.Println("Ошибка установки команд:", err)
	}
}
