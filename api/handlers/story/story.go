package hstory

import (
	dstory "github.com/rice9547/hakka_story/domain/story"
	sstory "github.com/rice9547/hakka_story/service/story"
)

type (
	Story struct {
		service sstory.Service
	}

	AudioResponse struct {
		ID       uint64 `json:"id"`
		AudioURL string `json:"audio_url"`
		Dialect  string `json:"dialect"`
	}

	PageResponse struct {
		ID           uint64          `json:"id"`
		Number       int             `json:"page_number"`
		ContentCN    string          `json:"content_cn"`
		ContentHakka string          `json:"content_hakka"`
		Audios       []AudioResponse `json:"audios"`
	}

	CategoryResponse struct {
		ID   uint64 `json:"id"`
		Name string `json:"name"`
	}

	StoryResponse struct {
		ID          uint64             `json:"id"`
		Title       string             `json:"title"`
		Description string             `json:"description"`
		CoverImage  string             `json:"cover_image"`
		Categories  []CategoryResponse `json:"categories"`
	}

	FullStoryResponse struct {
		StoryResponse
		Pages []PageResponse `json:"pages"`
	}
)

func New(service sstory.Service) *Story {
	return &Story{service: service}
}

func toResponse(story dstory.Story) StoryResponse {
	resp := StoryResponse{
		ID:          story.ID,
		Title:       story.Title,
		Description: story.Description,
	}

	if story.Image != nil {
		resp.CoverImage = story.Image.ImageURL
	}

	resp.Categories = make([]CategoryResponse, 0, len(story.Categories))
	for _, category := range story.Categories {
		resp.Categories = append(resp.Categories, CategoryResponse{
			ID:   category.ID,
			Name: category.Name,
		})
	}

	return resp
}

func toFullyResponse(story dstory.Story) FullStoryResponse {
	pages := make([]PageResponse, 0, len(story.Pages))
	for _, page := range story.Pages {
		pages = append(pages, PageResponse{
			ID:           page.ID,
			Number:       page.PageNumber,
			ContentCN:    page.ContentCN,
			ContentHakka: page.ContentHakka,
			Audios:       make([]AudioResponse, 0, len(page.AudioFiles)),
		})

		for _, audio := range page.AudioFiles {
			pages[len(pages)-1].Audios = append(pages[len(pages)-1].Audios, AudioResponse{
				ID:       audio.ID,
				AudioURL: audio.AudioURL,
				Dialect:  audio.Dialect,
			})
		}
	}

	return FullStoryResponse{
		StoryResponse: toResponse(story),
		Pages:         pages,
	}
}
