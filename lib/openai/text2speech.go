package openai

import (
	"context"
	"fmt"
)

const speechApiUrl = "https://api.openai.com/v1/audio/speech"

type AudioRequest struct {
	Model string `json:"model"`
	Input string `json:"input"`
	Voice string `json:"voice"`
}

func (c *Client) Text2Speech(ctx context.Context, prompt string) ([]byte, error) {
	requestData := AudioRequest{
		Model: "tts-1",
		Input: prompt,
		Voice: "alloy",
	}

	audioData, err := c.Post(ctx, speechApiUrl, requestData)
	if err != nil {
		return nil, fmt.Errorf("failed to generate audio: %v", err)
	}

	return audioData, nil
}
