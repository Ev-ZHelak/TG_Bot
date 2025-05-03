package main

import (
	"TG_Bot/config"
	"TG_Bot/internal/bot"
	"TG_Bot/internal/bot/handlers"
	"TG_Bot/internal/bot/menu"
	"TG_Bot/internal/db"
	"context"
	"fmt"
	"log"
)

func main() {
	fmt.Println(db.Pool)
	// Инициализация пула
	pool, err := db.InitPostgresPool(config.LoadPostgresConfig())
	if err != nil {
		log.Fatal(err)
	}
	db.Pool = pool
	defer pool.Close()

	// Работа с БД
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Release()

	var version string
	err = conn.QueryRow(context.Background(), "SELECT version()").Scan(&version)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("PostgreSQL version:", version)
	db.PrintStatistics(pool)
	fmt.Println(db.Pool)

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
