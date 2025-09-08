package bot

import (
	tb "gopkg.in/telebot.v3"

	"CS-SkinPulse/internal/storage"
)

func RegisterRoutes(b *tb.Bot, store storage.UserStore) {
	h := NewHandlers(store)

	b.Handle("/start", h.Start)
	b.Handle(h.btnProfile, h.ProfileInfo)
	b.Handle(h.btnInv, h.InventoryInfo)
	b.Handle(h.btnInvUpdate, h.UpdateInventory)
	b.Handle(h.btnInvStats, h.InventoryStats)
	b.Handle(h.btnBack, h.BackToMain)

	b.Handle(tb.OnText, h.OnText)
}
