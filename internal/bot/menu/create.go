package menu

import (
	"log"

	tele "gopkg.in/telebot.v3"
)

var (
	commands = []tele.Command{
		{Text: "start", Description: "Начать работу с ботом"},
		{Text: "help", Description: "Помощь и инструкции"},
		{Text: "settings", Description: "Настройки профиля"},
	}
)

// Устанавливаем список команд
func CreateMenu(b *tele.Bot) {
	err := b.SetCommands(commands)
	if err != nil {
		log.Println("Ошибка установки команд:", err)
	}
}
