package randomword_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tarusov/composer-challenge-bot/internal/config"
	randomword "github.com/tarusov/composer-challenge-bot/internal/random-word"
)

func TestClient(t *testing.T) {

	c := randomword.New(config.RandomWordConfig{
		APIURL: "https://random-word-api.vercel.app/api",
	})

	words, err := c.Words(3)
	if err != nil {
		t.Fatal(err)
	}
	assert.Len(t, words, 3)

	for _, word := range words {
		fmt.Println(word)
	}

}
