package storage

type UserStore interface {
	UpsertTelegramUser(u TelegramUser) error
	SaveSteamProfile(chatID int64, prof SteamProfile) error
	GetSteamProfile(chatID int64) (SteamProfile, bool)
}
