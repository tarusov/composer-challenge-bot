package generator_test

import (
	"fmt"

	"github.com/tarusov/composer-challenge-bot/internal/config"
	"github.com/tarusov/composer-challenge-bot/internal/generator"
)

type mockRandomWordAPI struct{}

func (m *mockRandomWordAPI) Words(count int) ([]string, error) {

	result := make([]string, 0, count)
	for i := 0; i < count; i++ {
		result = append(result, fmt.Sprintf("teapot_%d", i))
	}
	return result, nil
}

func mkTestGenerator() (*generator.Generator, error) {

	cfg, err := config.Read("../config_test.json")
	if err != nil {
		return nil, err
	}

	gen := generator.New(
		&mockRandomWordAPI{},
		cfg.Dictonary,
	)

	return gen, nil
}
