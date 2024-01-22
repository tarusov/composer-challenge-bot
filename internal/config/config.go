package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type (
	// TelegramConfig
	TelegramConfig struct {
		Token string `json:"token"`
		Debug bool   `json:"debug"`
	}

	// RandomWordConfig
	RandomWordConfig struct {
		APIURL string `json:"api_url"`
	}

	// Config
	Config struct {
		Telegram   TelegramConfig   `json:"telegram"`
		RandomWord RandomWordConfig `json:"random_word"`
		Dictonary  Dictonary        `json:"dict"`
	}

	Dictonary struct {
		TextHelloFn   []string `json:"text_hello_fn"`
		TextHelloFnLn []string `json:"text_hello_fn_ln"`
		TextHelloUn   []string `json:"text_hello_un"`
		TextKeyScale  []string `json:"text_key_scale"`
		Keys          []string `json:"keys"`
		Scales        []string `json:"scales"`
		TextTopics    []string `json:"text_topics"`
		TextGenre     []string `json:"text_genre"`
		GenreBase     []string `json:"genre_base"`
		GenrePrefix   []string `json:"genre_prefix"`
		GenrePatch    []string `json:"genre_patch"`
		Tips          []string `json:"tips"`
	}
)

// Read config from file.
func Read(fileName string) (*Config, error) {

	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	c := &Config{}
	err = json.Unmarshal(data, c)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config from json: %w", err)
	}

	return c, nil
}
