package storage

import (
	"errors"

	"gorm.io/gorm"
)

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

func (s *GormStore) GetTelegramUserByChatID(chatID int64) (TelegramUser, bool, error) {
	var u TelegramUser
	err := s.db.Where("chat_id = ?", chatID).First(&u).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return TelegramUser{}, false, nil
		}
		return TelegramUser{}, false, err
	}
	return u, true, nil
}

func (s *GormStore) SaveSteamProfile(telegramUserID uint, prof SteamProfile) error {
	prof.TelegramUserID = telegramUserID
	return s.db.Save(&prof).Error
}

func (s *GormStore) GetSteamProfileByChatID(chatID int64) (SteamProfile, bool) {
	var p SteamProfile
	err := s.db.
		Table("steam_profiles").
		Joins("JOIN telegram_users tu ON tu.id = steam_profiles.telegram_user_id").
		Where("tu.chat_id = ?", chatID).
		First(&p).Error
	if err != nil {
		return SteamProfile{}, false
	}
	return p, true
}
