package storage

type UserStore interface {
    UpsertTelegramUser(u TelegramUser) error
    GetTelegramUserByChatID(chatID int64) (TelegramUser, bool, error)

    SaveSteamProfile(telegramUserID uint, prof SteamProfile) error
    GetSteamProfileByChatID(chatID int64) (SteamProfile, bool)
}
