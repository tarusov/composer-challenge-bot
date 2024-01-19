package rw

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// aux client params.
const queryParamWordCount = "words"

type (
	client struct {
		url string
	}
)

// CTOR for random words api client.
func New(url string) *client {
	return &client{
		url: url,
	}
}

// Words return list of random words.
func (c *client) Words(ctx context.Context, count int) ([]string, error) {

	if count < 1 {
		count = 1
	}

	reqURL := fmt.Sprintf("%s?%s=%d", c.url, queryParamWordCount, count)

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL, nil)
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
