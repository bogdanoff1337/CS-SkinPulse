package bot

import tb "gopkg.in/telebot.v3"

func (h *Handlers) InventoryInfo(c tb.Context) error {
	if !h.ensureProfile(c) {
		return nil
	}
	return c.Send("Inventory section: choose an action below 👇", h.invMenu())
}

func (h *Handlers) UpdateInventory(c tb.Context) error {
	if !h.ensureProfile(c) {
		return nil
	}
	// TODO: call your inventory refresh service
	return c.Send("Started inventory update… this may take a bit.", h.invMenu())
}

func (h *Handlers) InventoryStats(c tb.Context) error {
	if !h.ensureProfile(c) {
		return nil
	}
	// TODO: fetch real metrics
	stats := "- Items: ~0\n- Total est. value: $0\n- Last refresh: n/a"
	return c.Send("Inventory stats:\n"+stats, h.invMenu())
}

func (h *Handlers) BackToMain(c tb.Context) error {
	return c.Send("Back to the main menu.", h.mainMenu())
}
