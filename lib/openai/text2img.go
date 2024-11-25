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

func (c *Client) Text2Image(prompt string) (string, []byte, error) {
	requestData := ImageRequest{
		Prompt: prompt,
		N:      1,
		Size:   "1024x1024",
	}

	jsonData, _ := json.Marshal(requestData)
	req, err := http.NewRequest("POST", imageApiUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", nil, err
	}
	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return "", nil, fmt.Errorf("failed to generate text: %s", body)
	}

	var response ImageResponse
	if err = json.Unmarshal(body, &response); err != nil {
		return "", nil, err
	}

	if len(response.Data) == 0 {
		return "", nil, fmt.Errorf("no image url found")
	}

	imageData, err := c.downloadImage(response.Data[0].URL)
	if err != nil {
		return "", nil, err
	}

	return response.Data[0].URL, imageData, nil
}

func (c *Client) downloadImage(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("download image failed: %s", body)
	}

	return io.ReadAll(resp.Body)
}
