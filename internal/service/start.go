package service

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const tgStartCmd = "/start"

// handleStartCmd send greetings message to chat.
func (s *Service) handleStartCmd(ctx context.Context, upd tgbotapi.Update) {

	fmt.Println("get start cmd")

	reply := tgbotapi.NewMessage(upd.Message.Chat.ID, s.rGreetings(upd))
	if _, err := s.tgAPI.Send(reply); err != nil {
		log.Println(err)
	}
}

func (s *Service) rGreetings(upd tgbotapi.Update) string {

	n := rand.Intn(len(s.greetings))
	gs := s.greetings[n]

	gs = strings.ReplaceAll(gs, "{{.USERNAME}}", upd.Message.From.UserName)
	gs = strings.ReplaceAll(gs, "{{.FIRST_NAME}}", upd.Message.From.FirstName)
	gs = strings.ReplaceAll(gs, "{{.LAST_NAME}}", upd.Message.From.LastName)

	return gs
}
