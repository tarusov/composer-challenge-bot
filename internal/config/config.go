package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type (
	TelegramConfig struct {
		Token string `yaml:"token"`
		Debug bool   `yaml:"debug"`
	}

	RandomWordConfig struct {
		URL string `yaml:"url"`
	}

	Config struct {
		Telegram   TelegramConfig   `yaml:"telegram"`
		RandomWord RandomWordConfig `yaml:"randomWord"`
		Keys       []string         `yaml:"keys"`
		Scales     []string         `yaml:"scales"`
		Greetings  []string         `yaml:"greetings"`
	}
)

// Read config file
func Read(fn string) (*Config, error) {

	data, err := os.ReadFile(fn)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	c := Config{}
	err = yaml.Unmarshal(data, &c)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	return &c, nil
}
