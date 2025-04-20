package main

import (
	"TG_Bot/config"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(config.API_TOKEN)
	if err != nil {
		panic(err)
	}

	bot.Debug = true

	// Создаем новую структуру UpdateConfig со смещением 0. Смещения используются
	// чтобы Telegram знал, что мы обработали предыдущие значения и они нам
	// не нужны повторно.
	updateConfig := tgbotapi.NewUpdate(0)

	// Указываем Telegram, что мы можем ждать до 30 секунд при каждом запросе
	// обновлений. Это позволяет получать информацию так же быстро, как при частых
	// запросах, но без необходимости отправлять их так много.
	updateConfig.Timeout = 30

	// Начинаем опрос Telegram для получения обновлений.
	updates := bot.GetUpdatesChan(updateConfig)

	// Обрабатываем каждое обновление, получаемое от Telegram.
	for update := range updates {
		// Telegram может отправлять разные типы обновлений в зависимости от
		// действий вашего бота. Пока мы хотим обрабатывать только сообщения,
		// поэтому игнорируем другие типы обновлений.
		if update.Message == nil {
			continue
		}

		// Теперь, когда мы получили новое сообщение, можем сформировать ответ!
		// Возьмем Chat ID и текст из входящего сообщения и используем их
		// для создания нового сообщения.
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		// Указываем, что это сообщение является ответом на предыдущее.
		// Для любых других параметров, кроме Chat ID или текста,
		// нужно устанавливать поля в структуре `MessageConfig`.
		msg.ReplyToMessageID = update.Message.MessageID

		// Отправляем наше сообщение! Нам не важно само отправленное сообщение,
		// поэтому просто игнорируем результат.
		if _, err := bot.Send(msg); err != nil {
			// Паника - не лучший способ обработки ошибок. У Telegram могут быть
			// перебои в работе или сетевые проблемы, лучше повторить отправку
			// или обрабатывать ошибки более аккуратно.
			panic(err)
		}
	}
}
