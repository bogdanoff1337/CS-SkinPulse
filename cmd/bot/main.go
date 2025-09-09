package main

import (
	"log"
	"os"
	"time"

	"CS-SkinPulse/internal/bot"
	"CS-SkinPulse/internal/storage"

	"github.com/joho/godotenv"
	tb "gopkg.in/telebot.v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	_ = godotenv.Load()

	token := os.Getenv("TELEGRAM_TOKEN")
	if token == "" {
		log.Fatal("set TELEGRAM_TOKEN")
	}

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("set DATABASE_URL (e.g. host=postgres user=postgres password=pass dbname=csskinpulse port=5432 sslmode=disable TimeZone=UTC)")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	userStore := storage.NewGormStore(db)

	pref := tb.Settings{
		Token:  token,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	}
	b, err := tb.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}

	ui := bot.NewUI()
	h := bot.NewHandlers(userStore)
	bot.RegisterRoutes(b, h, ui)

	log.Println("telegram bot started")
	b.Start()
}
