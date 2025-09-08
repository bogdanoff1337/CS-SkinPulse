package bot

import (
	tb "gopkg.in/telebot.v3"

	"CS-SkinPulse/internal/storage"
)

func RegisterRoutes(b *tb.Bot, store storage.UserStore) {
	h := NewHandlers(store)

	b.Handle("/start", h.Start)

	b.Handle(h.btnProfile.Text, h.ProfileInfo)
	b.Handle(h.btnInv.Text, h.InventoryInfo)
	b.Handle(h.btnInvUpdate.Text, h.UpdateInventory)
	b.Handle(h.btnInvStats.Text, h.InventoryStats)
	b.Handle(h.btnBack.Text, h.BackToMain)

	b.Handle(tb.OnText, h.OnText)
}
