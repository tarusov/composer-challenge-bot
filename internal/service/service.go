package service

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type (
	// Service struct.
	Service struct {
		tgAPI *tgbotapi.BotAPI
		txGen textGenerator
	}

	// textGenerator required methods.
	textGenerator interface {
		Hello(fn, ln, un string) string
		KeyScale() string
		Genre() string
		Instruments() string
		Topics() string
		Tips() string
		Bye() string
	}
)

// CTOR.
func New(bot *tgbotapi.BotAPI, gen textGenerator) *Service {
	return &Service{
		tgAPI: bot,
		txGen: gen,
	}
}

func (s *Service) ListenAndServe(ctx context.Context) {

	// Create update config
	updCfg := tgbotapi.NewUpdate(0)
	updCfg.Timeout = 30

	// Create update chan
	updCh := s.tgAPI.GetUpdatesChan(updCfg)

	for {
		select {
		case upd := <-updCh:
			if upd.Message == nil {
				continue
			}
			s.handleUpdate(upd)

		case <-ctx.Done():
			log.Println("Service terminated")
		}
	}
}

// handleUpdate
func (s *Service) handleUpdate(upd tgbotapi.Update) {

	var replyText string

	switch upd.Message.Text {
	case "/start":
		replyText = s.txGen.Hello(
			upd.Message.From.FirstName,
			upd.Message.From.LastName,
			upd.Message.From.UserName,
		)
	case "/roll":
		replyText = s.Roll()
	default:
		return
	}

	reply := tgbotapi.NewMessage(upd.Message.Chat.ID, replyText)
	reply.ParseMode = tgbotapi.ModeHTML

	if _, err := s.tgAPI.Send(reply); err != nil {
		log.Println(err)
	}
}

func (s *Service) Roll() string {

	var sb strings.Builder
	if _, err := sb.WriteString("Here is your composer challenge:\n\n"); err != nil {
		log.Println(err)
	}

	if _, err := sb.WriteString(s.txGen.KeyScale() + " "); err != nil {
		log.Println(err)
	}

	if _, err := sb.WriteString(s.txGen.Genre() + " "); err != nil {
		log.Println(err)
	}

	tempo := (rand.Intn(8) + 11) * 10
	if _, err := sb.WriteString(fmt.Sprintf("Tempo is %d.\n\n", tempo)); err != nil {
		log.Println(err)
	}

	if _, err := sb.WriteString(s.txGen.Topics() + "\n\n"); err != nil {
		log.Println(err)
	}

	if _, err := sb.WriteString(s.txGen.Instruments() + "\n\n"); err != nil {
		log.Println(err)
	}

	if _, err := sb.WriteString(s.txGen.Tips() + "\n\n"); err != nil {
		log.Println(err)
	}

	if _, err := sb.WriteString(s.txGen.Bye()); err != nil {
		log.Println(err)
	}

	return sb.String()
}
