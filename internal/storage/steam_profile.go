package storage

import "time"

type SteamProfile struct {
	ID             uint   `gorm:"primaryKey"`
	TelegramUserID uint   `gorm:"uniqueIndex;not null"` // відповідає telegram_user_id
	RawURL         string `gorm:"type:text"`
	SteamID64      string `gorm:"type:varchar(32);index"`
	Vanity         string `gorm:"type:varchar(255);index"`
	PersonaName    string `gorm:"type:varchar(255)"`
	Avatar         string `gorm:"type:varchar(512)"`
	SyncedAt       *time.Time
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
