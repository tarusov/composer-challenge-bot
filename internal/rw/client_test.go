package rw_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/tarusov/composer-challenge-bot/internal/rw"
)

func TestRandomWordClient(t *testing.T) {

	t.Skip("dev purpose only")

	c := rw.New(
		"https://random-word-api.vercel.app/api",
	)

	result, err := c.Words(context.Background(), 2)
	if err != nil {
		t.Fatal(err)
	}

	if len(result) != 2 {
		t.Fatal("unexpected behaviour")
	}

	fmt.Println(result)
}
