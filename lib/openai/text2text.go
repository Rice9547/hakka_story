package openai

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

	jsonData, _ := json.Marshal(requestData)
	req, _ := http.NewRequestWithContext(ctx, "POST", textApiUrl, bytes.NewBuffer(jsonData))
	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to generate text: %s", body)
	}

	var response Response
	if err := json.Unmarshal(body, &response); err != nil {
		return "", err
	}

	if len(response.Choices) > 0 {
		return response.Choices[0].Message.Content, nil
	} else {
		return "", fmt.Errorf("no response found")
	}
}
