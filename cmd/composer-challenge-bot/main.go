package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/tarusov/composer-challenge-bot/internal/config"
	"github.com/tarusov/composer-challenge-bot/internal/generator"
	randomword "github.com/tarusov/composer-challenge-bot/internal/random-word"
	"github.com/tarusov/composer-challenge-bot/internal/service"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {

	cfgFn := "config.json"
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

	gen := generator.New(
		randomword.New(cfg.RandomWord),
		cfg.Dictonary,
	)

	svc := service.New(
		tgAPI,
		gen,
	)

	// Run service
	ctx, cancel := context.WithCancel(context.Background())
	go svc.ListenAndServe(ctx)

	doneCh := make(chan struct{})

	// Make signal watcher
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)
	go func() {
		for range sigCh {
			log.Println("Received an interrupt...")
			cancel()
			close(doneCh)
		}
	}()

	<-doneCh
}
