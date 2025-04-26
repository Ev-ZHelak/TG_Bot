package models

type User struct {
	ID                      int64  `json:"id"`                          // Уникальный идентификатор пользователя
	FirstName               string `json:"first_name"`                  // Имя пользователя
	LastName                string `json:"last_name"`                   // Фамилия пользователя (может отсутствовать)
	Username                string `json:"username"`                    // Юзернейм (например, @username, может отсутствовать)
	LanguageCode            string `json:"language_code"`               // Код языка пользователя (например, "ru", "en")
	IsBot                   bool   `json:"is_bot"`                      // Является ли пользователь ботом
	IsPremium               bool   `json:"is_premium"`                  // Имеет ли пользователь Telegram Premium
	AddedToAttachmentMenu   bool   `json:"added_to_attachment_menu"`    // Добавлен ли бот в меню вложений
	CanJoinGroups           bool   `json:"can_join_groups"`             // Может ли бот вступать в группы
	CanReadAllGroupMessages bool   `json:"can_read_all_group_messages"` // Может ли бот читать все сообщения в группах
	SupportsInlineQueries   bool   `json:"supports_inline_queries"`     // Поддерживает ли бот inline-запросы
}
