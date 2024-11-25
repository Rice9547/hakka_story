package openai

import "github.com/rice9547/hakka_story/config"

type Client struct {
	apiKey string
}

func New(conf config.OpenAI) *Client {
	return &Client{
		apiKey: conf.APIKey,
	}
}
