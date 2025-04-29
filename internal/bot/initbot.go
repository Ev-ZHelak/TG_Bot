package bot

import (
	"TG_Bot/config"
	"log"
	"time"

	tele "gopkg.in/telebot.v3"
)

func InitBot() *tele.Bot {
	// Загрузка конфигурации
	cfg := config.LoadTelegramConfig()

	// Инициализация бота
	b, err := tele.NewBot(tele.Settings{
		Token:   cfg.Telegram.Token,
		Verbose: cfg.Telegram.Debug,
		Poller:  &tele.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
	}
	return b
}
