package main

import (
	"TG_Bot/config"
	"TG_Bot/internal/bot"
	"TG_Bot/internal/bot/handlers"
	"TG_Bot/internal/bot/menu"
	"fmt"
)

func main() {
	fmt.Println(config.LoadPostgresConfig().Postgresql)
	// Создание бота
	b := bot.InitBot()
	// b.Use()
	// Создаем меню из списка команд
	menu.CreateMenu(b)
	// Основной обработчик команд
	handlers.MainHandlerCommands(b)
	// Основной обработчик сообщений
	handlers.MainHandlerMessages(b)
	fmt.Println("Start bot...")
	b.Start()
}
