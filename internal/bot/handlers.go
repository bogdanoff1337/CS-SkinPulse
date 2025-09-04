package bot

import (
    "fmt"
    "strings"

    tb "gopkg.in/telebot.v3"

    "CS-SkinPulse/internal/steam"
    "CS-SkinPulse/internal/storage"
)

type Handlers struct {
    store       storage.UserStore
    btnSendLink tb.Btn
    menu        *tb.ReplyMarkup
}

func NewHandlers(store storage.UserStore) *Handlers {
    menu := &tb.ReplyMarkup{ResizeKeyboard: true}
    btn := menu.Text("🔗 Send Steam Link")
    menu.Reply(menu.Row(btn))

    return &Handlers{
        store:       store,
        btnSendLink: btn,
        menu:        menu,
    }
}

func (h *Handlers) Start(c tb.Context) error {
    msg := "Hi! Please send me your steam account link, like:\n" +
        "• https://steamcommunity.com/id/<your_nick>\n" +
        "• or https://steamcommunity.com/profiles/<steamID64>\n\n"
    return c.Send(msg, h.menu)
}

func (h *Handlers) AskLink(c tb.Context) error {
    return c.Send("Send me your Steam profile link:", h.menu)
}

func (h *Handlers) Me(c tb.Context) error {
    chatID := c.Chat().ID
    if prof, ok := h.store.GetSteamProfile(chatID); ok {
        if prof.SteamID64 != "" {
            return c.Send(fmt.Sprintf("Твій профіль: %s\nsteamID64: `%s`", prof.RawURL, prof.SteamID64),
                &tb.SendOptions{ParseMode: tb.ModeMarkdown})
        }
        return c.Send("Your profile: " + prof.RawURL)
    }
    return c.Send("You haven't set your Steam profile yet. Use the button below to send it.", h.menu)
}

func (h *Handlers) OnText(c tb.Context) error {
    text := strings.TrimSpace(c.Text())
    if !steam.IsValidSteamURL(text) {
        return c.Send("Link is not correct" +
            "Try again (examples):\n" +
            "• https://steamcommunity.com/id/your_nickname\n" +
            "• https://steamcommunity.com/profiles/76561198000000000")
    }

    prof := storage.SteamProfile{RawURL: text}
    if id64, ok := steam.ExtractSteamID64(text); ok {
        prof.SteamID64 = id64
    }

    if err := h.store.SaveSteamProfile(c.Chat().ID, prof); err != nil {
        return c.Send("Error saving profile. Try again later.")
    }

    if prof.SteamID64 != "" {
        return c.Send("✅ Profile is save, your id **steamID64**: `"+prof.SteamID64+"`",
            &tb.SendOptions{ParseMode: tb.ModeMarkdown})
    }
    return c.Send("✅ Profile is save: " + prof.RawURL)
}
