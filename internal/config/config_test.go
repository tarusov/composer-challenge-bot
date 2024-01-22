package config_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tarusov/composer-challenge-bot/internal/config"
)

func TestConfigRead(t *testing.T) {

	cfg, err := config.Read("../config_test.json")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, cfg.Telegram.Token, "qqck")
	assert.Equal(t, cfg.Telegram.Debug, true)
	assert.Equal(t, cfg.RandomWord.APIURL, "http://rwapi")
	assert.Len(t, cfg.Dictonary.TextHelloFn, 3)
	assert.Len(t, cfg.Dictonary.TextHelloFnLn, 3)
	assert.Len(t, cfg.Dictonary.TextHelloUn, 3)
	assert.Len(t, cfg.Dictonary.TextKeyScale, 3)
	assert.Len(t, cfg.Dictonary.Keys, 12)
	assert.Len(t, cfg.Dictonary.Scales, 3)
	assert.Len(t, cfg.Dictonary.TextTopics, 3)
}
