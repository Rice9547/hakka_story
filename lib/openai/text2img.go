package openai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const imageApiUrl = "https://api.openai.com/v1/images/generations"

type ImageRequest struct {
	Prompt string `json:"prompt"`
	N      int    `json:"n"`
	Size   string `json:"size"`
}

type ImageResponse struct {
	Data []struct {
		URL string `json:"url"`
	} `json:"data"`
}

func (c *Client) Text2Image(prompt string) (string, error) {
	requestData := ImageRequest{
		Prompt: prompt,
		N:      1,
		Size:   "1024x1024",
	}

	jsonData, _ := json.Marshal(requestData)
	req, err := http.NewRequest("POST", imageApiUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to generate text: %s", body)
	}

	var response ImageResponse
	if err = json.Unmarshal(body, &response); err != nil {
		return "", err
	}

	if len(response.Data) > 0 {
		return response.Data[0].URL, nil
	} else {
		return "", fmt.Errorf("no image url found")
	}
}
