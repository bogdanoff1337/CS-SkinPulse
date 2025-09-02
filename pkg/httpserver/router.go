package httpserver

import (
	"CS-SkinPulse/internal/steam"
	"net/http"

	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func Router(corsAllowed []string, steamClient *steam.Client) http.Handler {
	r := chi.NewRouter()
	r.Use(chimw.RealIP, chimw.Logger, chimw.Recoverer)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   corsAllowed,
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.Post("/api/price", steam.PriceHandler(steamClient))

	return r
}
