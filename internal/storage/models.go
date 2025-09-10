package storage

import "time"

type TelegramUser struct {
    ID           uint   `gorm:"primaryKey"`
    ChatID       int64  `gorm:"uniqueIndex;not null"`
    Username     string `gorm:"index"`
    FirstName    string
    LastName     string
    LanguageCode string
    Timezone     string `gorm:"default:Europe/Kyiv"`
    CreatedAt    time.Time
    UpdatedAt    time.Time
}
