package bot

import (
	tb "gopkg.in/telebot.v3"

	"CS-SkinPulse/internal/storage"
)

type Handlers struct {
	store storage.UserStore
	ui    *UI
}

func NewHandlers(store storage.UserStore, ui *UI) *Handlers {
	return &Handlers{store: store, ui: ui}
}

func (h *Handlers) mainMenu() *tb.ReplyMarkup { return h.ui.Main }
func (h *Handlers) invMenu() *tb.ReplyMarkup  { return h.ui.InvMenu }
