package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"CS-SkinPulse/internal/steam"
	"CS-SkinPulse/pkg/config"
	"CS-SkinPulse/pkg/httpserver"
)

func main() {
	cfg := config.Load()
	steamClient := steam.NewClient()
	router := httpserver.Router(cfg.CORSAllowed, steamClient)

	srv := &http.Server{
		Addr:              cfg.Addr,
		Handler:           router,
		ReadHeaderTimeout: 10 * time.Second,
	}

	go func() {
		log.Printf("env=%s listen=%s", cfg.Env, cfg.Addr)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal("http serve:", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_ = srv.Shutdown(ctx)
	log.Println("👋 server stopped")
}
