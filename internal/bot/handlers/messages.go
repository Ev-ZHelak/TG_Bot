package handlers

import (
	"fmt"

	tele "gopkg.in/telebot.v3"
)

// Обработчики команд
func MainHandlerMessages(b *tele.Bot) {
	b.Handle(tele.OnText, func(c tele.Context) error {
		user := c.Sender() // Получаем *User отправителя
		c.Send(fmt.Sprintf(
			"Ваш ID: %d\nИмя: %s\nЮзернейм: @%s",
			user.ID,
			user.FirstName,
			user.Username,
		))
		return nil
	})
}
