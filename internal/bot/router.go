package bot

import (
    tb "gopkg.in/telebot.v3"

    "CS-SkinPulse/internal/storage"
)

func RegisterRoutes(b *tb.Bot, store storage.UserStore) {
    h := NewHandlers(store)

    b.Handle("/start", h.Start)
    b.Handle("/me", h.Me)

    b.Handle(&h.btnSendLink, h.AskLink)

    b.Handle(tb.OnText, h.OnText)
}
