package randomword

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/tarusov/composer-challenge-bot/internal/config"
)

const queryParamWordCount = "words"

type (
	// Client for random word API.
	Client struct {
		apiURL string
	}
)

// CTOR
func New(cfg config.RandomWordConfig) *Client {
	return &Client{
		apiURL: cfg.APIURL,
	}
}

// Words return list of random words.
func (c *Client) Words(count int) ([]string, error) {

	if count < 1 {
		count = 1
	}

	reqURL := fmt.Sprintf("%s?%s=%d", c.apiURL, queryParamWordCount, count)

	httpReq, err := http.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create http request: %w", err)
	}

	httpResp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("failed to execute http request: %w", err)
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected response status code: %d", httpResp.StatusCode)
	}

	respBody, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var result = make([]string, 0, count)
	err = json.Unmarshal(respBody, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to parse response body: %w", err)
	}

	return result, nil
}
