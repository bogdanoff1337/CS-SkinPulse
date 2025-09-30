package bot

import (
	"fmt"
	"strings"

	tb "gopkg.in/telebot.v3"

	"CS-SkinPulse/internal/steam"
	"CS-SkinPulse/internal/storage"
)

type Handlers struct {
	store storage.UserStore

	mainMenu   *tb.ReplyMarkup
	btnProfile tb.Btn
	btnInv     tb.Btn

	menuInv      *tb.ReplyMarkup
	btnInvUpdate tb.Btn
	btnInvStats  tb.Btn
	btnBack      tb.Btn
}

func NewHandlers(store storage.UserStore) *Handlers {
	main := &tb.ReplyMarkup{ResizeKeyboard: true}
	btnProfile := main.Text("👤 My profile info")
	btnInv := main.Text("📦 Inventory info")
	main.Reply(main.Row(btnProfile, btnInv))

	inv := &tb.ReplyMarkup{ResizeKeyboard: true}
	btnInvUpdate := inv.Text("🔄 Update inventory")
	btnInvStats := inv.Text("📊 Inventory stats")
	btnBack := inv.Text("⬅️ Back")
	inv.Reply(inv.Row(btnInvUpdate, btnInvStats), inv.Row(btnBack))

	return &Handlers{
		store:        store,
		mainMenu:     main,
		btnProfile:   btnProfile,
		btnInv:       btnInv,
		menuInv:      inv,
		btnInvUpdate: btnInvUpdate,
		btnInvStats:  btnInvStats,
		btnBack:      btnBack,
	}
}

func (h *Handlers) Start(c tb.Context) error {
	_ = h.store.UpsertTelegramUser(storage.TelegramUser{
		ChatID:       c.Chat().ID,
		Username:     c.Sender().Username,
		FirstName:    c.Sender().FirstName,
		LastName:     c.Sender().LastName,
		LanguageCode: c.Sender().LanguageCode,
		Timezone:     "Europe/Kyiv",
	})

	msg := "Hi! Send me your Steam profile once, then use the menu below.\n\n" +
		"Examples:\n" +
		"• https://steamcommunity.com/id/<your_nick>\n" +
		"• https://steamcommunity.com/profiles/<steamID64>\n"
	return c.Send(msg, h.mainMenu)
}

func (h *Handlers) ProfileInfo(c tb.Context) error {
	prof, ok := h.store.GetSteamProfileByChatID(c.Chat().ID)
	if ok && (prof.SteamID64 != "" || prof.RawURL != "") {
		if prof.SteamID64 != "" {
			return c.Send(
				fmt.Sprintf("Your profile: %s\nsteamID64: `%s`", prof.RawURL, prof.SteamID64),
				&tb.SendOptions{ParseMode: tb.ModeMarkdown, ReplyMarkup: h.mainMenu},
			)
		}
		return c.Send("Your profile: "+prof.RawURL, h.mainMenu)
	}
	return c.Send(
		"I don't see your Steam link yet. Please send your profile URL as a message.\n\n"+
			"• https://steamcommunity.com/id/your_nickname\n"+
			"• https://steamcommunity.com/profiles/76561198000000000",
		h.mainMenu,
	)
}

func (h *Handlers) InventoryInfo(c tb.Context) error {
	if !h.ensureProfile(c) {
		return nil
	}
	return c.Send("Inventory section: choose an action below 👇", h.menuInv)
}

func (h *Handlers) UpdateInventory(c tb.Context) error {
	if !h.ensureProfile(c) {
		return nil
	}
	// TODO: call your inventory refresh service
	return c.Send("Started inventory update… this may take a bit.", h.menuInv)
}

func (h *Handlers) InventoryStats(c tb.Context) error {
	if !h.ensureProfile(c) {
		return nil
	}
	// TODO: fetch real metrics from store/DB
	stats := "- Items: ~0\n- Total est. value: $0\n- Last refresh: n/a"
	return c.Send("Inventory stats:\n"+stats, h.menuInv)
}

func (h *Handlers) BackToMain(c tb.Context) error {
	return c.Send("Back to the main menu.", h.mainMenu)
}

func (h *Handlers) OnText(c tb.Context) error {
	text := strings.TrimSpace(c.Text())
	if !steam.IsValidSteamURL(text) {
		return c.Send("Invalid link.\nTry again:\n• https://steamcommunity.com/id/your_nickname\n• https://steamcommunity.com/profiles/76561198000000000", h.mainMenu)
	}

	u, ok, err := h.store.GetTelegramUserByChatID(c.Chat().ID)
	if err != nil {
		return c.Send("DB error. Try later.", h.mainMenu)
	}
	if !ok {
		if err := h.store.UpsertTelegramUser(storage.TelegramUser{
			ChatID:       c.Chat().ID,
			Username:     c.Sender().Username,
			FirstName:    c.Sender().FirstName,
			LastName:     c.Sender().LastName,
			LanguageCode: c.Sender().LanguageCode,
			Timezone:     "Europe/Kyiv",
		}); err != nil {
			return c.Send("Failed to register you. Try /start again.", h.mainMenu)
		}
		u, _, _ = h.store.GetTelegramUserByChatID(c.Chat().ID)
	}

	prof := storage.SteamProfile{RawURL: text}
	if id64, ok := steam.ExtractSteamID64(text); ok {
		prof.SteamID64 = id64
	}

	if err := h.store.SaveSteamProfile(u.ID, prof); err != nil {
		return c.Send("Failed to save profile. Please try again later.", h.mainMenu)
	}

	if prof.SteamID64 != "" {
		return c.Send("✅ Profile saved. Your **steamID64**: `"+prof.SteamID64+"`",
			&tb.SendOptions{ParseMode: tb.ModeMarkdown, ReplyMarkup: h.mainMenu})
	}
	return c.Send("✅ Profile saved: "+prof.RawURL, h.mainMenu)
}

func (h *Handlers) ensureProfile(c tb.Context) bool {
	prof, ok := h.store.GetSteamProfileByChatID(c.Chat().ID)
	if ok && (prof.SteamID64 != "" || prof.RawURL != "") {
		return true
	}
	_ = c.Send("Please send your Steam profile (URL) first.", h.mainMenu)
	return false
}

func (h *Handlers) MainMenu() *tb.ReplyMarkup {
	return h.mainMenu
}
