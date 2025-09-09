package storage

type SteamProfile struct {
	ChatID    int64  `gorm:"primaryKey;not null"`
	RawURL    string `gorm:"type:text;not null"`
	SteamID64 string `gorm:"type:varchar(32);index"`
}
