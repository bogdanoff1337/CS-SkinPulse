package bot

import tb "gopkg.in/telebot.v3"

type UI struct {
	Main       *tb.ReplyMarkup
	BtnProfile tb.Btn
	BtnInv     tb.Btn

	InvMenu      *tb.ReplyMarkup
	BtnInvInsert tb.Btn
	BtnInvUpdate tb.Btn
	BtnInvStats  tb.Btn
	BtnBack      tb.Btn
}

func NewUI() *UI {
	main := &tb.ReplyMarkup{ResizeKeyboard: true}
	btnProfile := main.Text("👤 My profile info")
	btnInv := main.Text("📦 Inventory info")
	main.Reply(main.Row(btnProfile, btnInv))

	inv := &tb.ReplyMarkup{ResizeKeyboard: true}
	btnInvLoad := inv.Text("📥 Load")
	btnInvUpdate := inv.Text("🔄 Update")
	btnInvStats := inv.Text("📊 Stats")
	btnBack := inv.Text("⬅️ Back")
	inv.Reply(inv.Row(btnInvLoad, btnInvUpdate, btnInvStats), inv.Row(btnBack))

	return &UI{
		Main:         main,
		BtnProfile:   btnProfile,
		BtnInv:       btnInv,
		InvMenu:      inv,
		BtnInvInsert: btnInvLoad,
		BtnInvUpdate: btnInvUpdate,
		BtnInvStats:  btnInvStats,
		BtnBack:      btnBack,
	}
}
