package openai

import (
	"github.com/rice9547/hakka_story/config"
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
