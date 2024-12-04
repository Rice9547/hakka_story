package openai

import (
	"context"
	"encoding/json"
	"fmt"
)

const textApiUrl = "https://api.openai.com/v1/chat/completions"

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Request struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type Response struct {
	Choices []struct {
		Message Message `json:"message"`
	} `json:"choices"`
}

func (c *Client) Text2Text(ctx context.Context, prompt string) (string, error) {
	messages := []Message{
		{Role: "system", Content: "你是一個客語專家，接下來你會收到一些中文的故事段落，請翻譯成四縣腔的客語，不需要額外的回覆。"},
		{Role: "user", Content: prompt},
	}

	requestData := Request{
		Model:    "gpt-4",
		Messages: messages,
	}

	respData, err := c.Post(ctx, textApiUrl, requestData)
	if err != nil {
		return "", fmt.Errorf("failed to generate text: %v", err)
	}

	response := Response{}
	if err := json.Unmarshal(respData, &response); err != nil {
		return "", fmt.Errorf("failed to unmarshal response: %v", err)
	}

	if len(response.Choices) > 0 {
		return response.Choices[0].Message.Content, nil
	}

	return "", fmt.Errorf("no response found")
}
