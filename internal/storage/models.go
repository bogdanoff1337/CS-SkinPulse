package storage

type TelegramUser struct {
	ID           uint  `gorm:"primaryKey"`
	ChatID       int64 `gorm:"uniqueIndex;not null"`
	Username     string
	FirstName    string
	LastName     string
	LanguageCode string
	Timezone     string
}
