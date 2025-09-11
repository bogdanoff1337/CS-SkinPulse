package bot

import (
	"fmt"
	"strings"

	tb "gopkg.in/telebot.v3"

	"CS-SkinPulse/internal/steam"
	"CS-SkinPulse/internal/storage"
)

func (h *Handlers) Start(c tb.Context) error {
	msg := "Hi! Send me your Steam profile once, then use the menu below.\n\n" +
		"Examples:\n" +
		"• https://steamcommunity.com/id/<your_nick>\n" +
		"• https://steamcommunity.com/profiles/<steamID64>\n"
	return c.Send(msg, h.mainMenu())
}

func (h *Handlers) ProfileInfo(c tb.Context) error {
	chatID := c.Chat().ID
	if prof, ok := h.store.GetSteamProfile(chatID); ok && (prof.SteamID64 != "" || prof.RawURL != "") {
		if prof.SteamID64 != "" {
			return c.Send(
				fmt.Sprintf("Your profile: %s\nsteamID64: `%s`", prof.RawURL, prof.SteamID64),
				&tb.SendOptions{ParseMode: tb.ModeMarkdown, ReplyMarkup: h.mainMenu()},
			)
		}
		return c.Send("Your profile: "+prof.RawURL, h.mainMenu())
	}
	return c.Send(
		"I don't see your Steam link yet. Please send your profile URL as a message.\n\n"+
			"• https://steamcommunity.com/id/your_nickname\n"+
			"• https://steamcommunity.com/profiles/76561198000000000",
		h.mainMenu(),
	)
}

func (h *Handlers) SaveProfileFromMessage(c tb.Context) error {
	text := strings.TrimSpace(c.Text())

	if !steam.IsValidSteamURL(text) {
		return c.Send(
			"Invalid link.\n"+
				"Try again (examples):\n"+
				"• https://steamcommunity.com/id/your_nickname\n"+
				"• https://steamcommunity.com/profiles/76561198000000000",
			h.mainMenu(),
		)
	}

	prof := storage.SteamProfile{RawURL: text}
	if id64, ok := steam.ExtractSteamID64(text); ok {
		prof.SteamID64 = id64
	}
	//
	//if err := h.store.SaveSteamProfile(c.Chat().ID, prof); err != nil {
	//	return c.Send("Failed to save profile. Please try again later.", h.mainMenu())
	//}

	if prof.SteamID64 != "" {
		return c.Send(
			"✅ Profile saved. Your **steamID64**: `"+prof.SteamID64+"`",
			&tb.SendOptions{ParseMode: tb.ModeMarkdown, ReplyMarkup: h.mainMenu()},
		)
	}
	return c.Send("✅ Profile saved: "+prof.RawURL, h.mainMenu())
}

func (h *Handlers) ensureProfile(c tb.Context) bool {
	chatID := c.Chat().ID
	if prof, ok := h.store.GetSteamProfile(chatID); ok && (prof.SteamID64 != "" || prof.RawURL != "") {
		return true
	}
	_ = c.Send("Please send your Steam profile (URL) first.", h.mainMenu())
	return false
}
