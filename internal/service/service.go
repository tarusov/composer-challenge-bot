package service

import (
	"context"
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type (
	// Service is main structure of bot.
	Service struct {
		tgAPI          *tgbotapi.BotAPI
		updateInterval time.Duration
		rwAPI          randomWordAPI
		greetings      []string
		keys           []string
		scales         []string
		instruments    []string
		genres         []string
	}

	// randomWordAPI define a methods for random words api.
	randomWordAPI interface {
		Words(context.Context, int) ([]string, error)
	}
)

// New create new service instance.
func New(tgAPI *tgbotapi.BotAPI, rwAPI randomWordAPI, opts ...serviceOption) *Service {

	var s = &Service{
		tgAPI:          tgAPI,
		updateInterval: time.Second * 60,
		rwAPI:          rwAPI,
		keys: []string{
			"C",
		},
		scales: []string{
			"Major",
		},
		genres: []string{
			"rock",
		},
		instruments: []string{
			"guitar",
			"bass",
			"drums",
		},
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

// ListenAndServe handle requests.
func (s *Service) ListenAndServe(ctx context.Context) {

	updCfg := tgbotapi.NewUpdate(0)
	updCfg.Timeout = int(s.updateInterval.Seconds())
	updCh := s.tgAPI.GetUpdatesChan(updCfg)

	for {
		select {
		case update := <-updCh:

			if update.Message == nil {
				continue
			}

			switch {
			case update.Message.Text == tgStartCmd:
				s.handleStartCmd(ctx, update)

			case update.Message.Text == tgRollCmd:
				s.handleRollCmd(ctx, update)
			}

		case <-ctx.Done():
			log.Println("Service stopped...")
			return
		}

	}
}
