package storage

type SteamProfile struct {
	RawURL    string
	SteamID64 string
}

type UserStore interface {
	SaveSteamProfile(chatID int64, prof SteamProfile) error
	GetSteamProfile(chatID int64) (SteamProfile, bool)
}

type MemoryStore struct {
	data map[int64]SteamProfile
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{data: make(map[int64]SteamProfile)}
}

func (s *MemoryStore) SaveSteamProfile(chatID int64, prof SteamProfile) error {
	s.data[chatID] = prof
	return nil
}

func (s *MemoryStore) GetSteamProfile(chatID int64) (SteamProfile, bool) {
	prof, ok := s.data[chatID]
	return prof, ok
}
