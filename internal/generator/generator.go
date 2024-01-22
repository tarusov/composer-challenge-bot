package generator

import (
	"math/rand"

	"github.com/tarusov/composer-challenge-bot/internal/config"
)

type (
	// Generator struct.
	Generator struct {
		rwAPI randomWordAPI
		dict  config.Dictonary
	}

	// randomWordAPI define methods for random word generator.
	randomWordAPI interface {
		Words(int) ([]string, error)
	}
)

// CTOR
func New(rwAPI randomWordAPI, dict config.Dictonary) *Generator {
	return &Generator{
		rwAPI: rwAPI,
		dict:  dict,
	}
}

// randomElem
func randomElem(list []string) string {
	if l := len(list); l > 0 {
		return list[rand.Intn(l)]
	}
	return ""
}
