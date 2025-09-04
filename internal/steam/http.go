package steam

import (
	"encoding/json"
	"net/http"
)

type PriceRequest struct {
	Item string `json:"item"`
}

func PriceHandler(client *Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		var req PriceRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}
		price, err := client.FetchPrice(req.Item, 1)
		if err != nil {
			http.Error(w, "Failed to fetch price", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(price)
	}
}
