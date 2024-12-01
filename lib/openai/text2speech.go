package openai

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const speechApiUrl = "https://api.openai.com/v1/audio/speech"

type AudioRequest struct {
	Model string `json:"model"`
	Input string `json:"input"`
	Voice string `json:"voice"`
}

type AudioResponse struct {
	AudioContent string `json:"audio_content"`
}

func (c *Client) Text2Speech(ctx context.Context, prompt string) ([]byte, error) {
	requestData := AudioRequest{
		Model: "tts-1",
		Input: prompt,
		Voice: "alloy",
	}

	jsonData, _ := json.Marshal(requestData)

	req, _ := http.NewRequestWithContext(ctx, "POST", speechApiUrl, bytes.NewBuffer(jsonData))
	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("generate audio failed: %s", body)
	}

	audioData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read audio data failed: %v", err)
	}

	return audioData, nil
}
