package main

import (
	"TG_Bot/internal/bot/handlers"
	"TG_Bot/internal/bot/menu"
	"TG_Bot/internal/config"
	"fmt"
	"log"
	"time"

	tele "gopkg.in/telebot.v3"
)

func main() {
	// Загрузка конфигурации
	cfg := config.LoadTgBotConfig()

	// Инициализация бота
	b, err := tele.NewBot(tele.Settings{
		Token:   cfg.Telegram.Token,
		Verbose: cfg.Telegram.Debug,
		Poller:  &tele.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
	}

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
