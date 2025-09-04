package main

import (
    "log"
    "os"
    "time"

    "CS-SkinPulse/internal/bot"
    "CS-SkinPulse/internal/storage"
    "github.com/joho/godotenv"
    tb "gopkg.in/telebot.v3"
)

func main() {
    _ = godotenv.Load()

    token := os.Getenv("TELEGRAM_TOKEN")
    if token == "" {
        log.Fatal("set TELEGRAM_TOKEN env var")
    }

    userStore := storage.NewMemoryStore()

    pref := tb.Settings{
        Token:  token,
        Poller: &tb.LongPoller{Timeout: 10 * time.Second},
    }

    tg, err := tb.NewBot(pref)
    if err != nil {
        log.Fatal(err)
    }

    bot.RegisterRoutes(tg, userStore)

    log.Println("telegram bot started")
    tg.Start()
}
