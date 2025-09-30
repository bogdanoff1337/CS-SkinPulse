package bot

import tb "gopkg.in/telebot.v3"

func RegisterRoutes(b *tb.Bot, h *Handlers, ui *UI) {
	b.Handle("/start", h.Start)

	b.Handle(ui.BtnProfile.Text, h.ProfileInfo)
	b.Handle(ui.BtnInv.Text, h.InventoryInfo)
	b.Handle(ui.BtnInvInsert.Text, h.LoadInventory)
	b.Handle(ui.BtnInvUpdate.Text, h.UpdateInventory)
	b.Handle(ui.BtnInvStats.Text, h.InventoryStats)
	b.Handle(ui.BtnBack.Text, h.BackToMain)

	b.Handle(tb.OnText, h.OnText)
}
