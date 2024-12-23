package openai

import (
	"context"
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

func (c *Client) Text2Image(ctx context.Context, prompt string) (string, []byte, error) {
	requestData := ImageRequest{
		Prompt: prompt,
		N:      1,
		Size:   "1024x1024",
	}

	respData, err := c.Post(ctx, imageApiUrl, requestData)
	if err != nil {
		return "", nil, fmt.Errorf("failed to generate image: %v", err)
	}

	response := ImageResponse{}
	if err := json.Unmarshal(respData, &response); err != nil {
		return "", nil, fmt.Errorf("failed to unmarshal response: %v", err)
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
