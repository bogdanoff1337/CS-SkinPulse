package storage

import "gorm.io/gorm"

type GormStore struct{ db *gorm.DB }

func NewGormStore(db *gorm.DB) *GormStore {
	return &GormStore{db: db}
}

func (s *GormStore) UpsertTelegramUser(u TelegramUser) error {
	var existing TelegramUser
	if err := s.db.Where("chat_id = ?", u.ChatID).First(&existing).Error; err == nil {
		u.ID = existing.ID
	}
	return s.db.Save(&u).Error
}

func (s *GormStore) SaveSteamProfile(chatID int64, prof SteamProfile) error {
	prof.ChatID = chatID
	return s.db.Save(&prof).Error
}

func (s *GormStore) GetSteamProfile(chatID int64) (SteamProfile, bool) {
	var p SteamProfile
	if err := s.db.First(&p, "telegram_user_id = ?", chatID).Error; err != nil {
		return SteamProfile{}, false
	}
	return p, true
}
