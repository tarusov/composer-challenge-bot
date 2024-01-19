package service

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const tgRollCmd = "/roll"

func (s *Service) handleRollCmd(ctx context.Context, upd tgbotapi.Update) {

	b := strings.Builder{}

	// Greetings.

	//
	b.WriteString(fmt.Sprintf("How about to write a song in <b>%s</b> key in <b>%s</b> scale?\n", s.rKey(), s.rScale()))

	b.WriteString("And what can this track about:")
	topics := s.rTopics(ctx)
	for _, t := range topics {
		_, _ = b.WriteString(t + " ")
	}

	reply := tgbotapi.NewMessage(upd.Message.Chat.ID, b.String())
	if _, err := s.tgAPI.Send(reply); err != nil {
		log.Println(err)
	}
}

func (s *Service) rKey() string {
	n := rand.Intn(len(s.keys))
	return s.keys[n]
}

func (s *Service) rScale() string {
	n := rand.Intn(len(s.scales))
	return s.scales[n]
}

func (s *Service) rTopics(ctx context.Context) []string {
	resp, err := s.rwAPI.Words(ctx, rand.Intn(3)+2)
	if err != nil {
		return []string{"peace", "loving", "understanding"}
	}
	return resp
}
