package openai

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/rice9547/hakka_story/config"
	"io"
	"net/http"
)

type Client struct {
	apiKey     string
	httpClient *http.Client
}

func New(conf config.OpenAI) *Client {
	return &Client{
		apiKey:     conf.APIKey,
		httpClient: &http.Client{},
	}
}

func (c *Client) Post(ctx context.Context, uri string, reqJson any) ([]byte, error) {
	jsonData, err := json.Marshal(reqJson)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request json: %v", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", uri, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create post request: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to do post request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to do post request with status code: %d, parse body err: %v", resp.StatusCode, err)
		}
		return nil, fmt.Errorf("failed to do post request with status code: %d, err: %v", resp.StatusCode, body)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	return respBody, nil
}
