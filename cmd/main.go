package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/tarusov/composer-challenge-bot/internal/config"
	"github.com/tarusov/composer-challenge-bot/internal/rw"
	"github.com/tarusov/composer-challenge-bot/internal/service"
)

func main() {

	cfgFn := "real_config.yml"
	if len(os.Args) != 1 {
		cfgFn = os.Args[1]
	}

	cfg, err := config.Read(cfgFn)
	if err != nil {
		log.Fatal("read config", err)
	}

	tgAPI, err := tgbotapi.NewBotAPI(cfg.Telegram.Token)
	if err != nil {
		log.Fatal("connect telegram", err)
	}
	tgAPI.Debug = cfg.Telegram.Debug

	service := service.New(
		tgAPI,
		rw.New(cfg.RandomWord.URL),
		service.WithGreetings(cfg.Greetings),
		service.WithKeys(cfg.Keys),
		service.WithScales(cfg.Scales),
	)

	ctx, cancel := context.WithCancel(context.Background())
	go service.ListenAndServe(ctx)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)
	for range sigCh {
		log.Printf("\nReceived an interrupt...\n\n")

		// Dirty way to stop service. Sorry, i don't give a f...
		cancel()
		time.Sleep(time.Millisecond * 100)
		break
	}
}
