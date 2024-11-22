package hstory

import (
	dstory "github.com/rice9547/hakka_story/domain/story"
	sstory "github.com/rice9547/hakka_story/service/story"
)

type (
	Story struct {
		service sstory.Service
	}

	PageResponse struct {
		ID           uint64 `json:"id"`
		Number       int    `json:"page_number"`
		ContentCN    string `json:"content_cn"`
		ContentHakka string `json:"content_hakka"`
	}

	StoryResponse struct {
		ID          uint64 `json:"id"`
		Title       string `json:"title"`
		Description string `json:"description"`
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
	return StoryResponse{
		ID:          story.ID,
		Title:       story.Title,
		Description: story.Description,
	}
}

func toFullyResponse(story dstory.Story) FullStoryResponse {
	pages := make([]PageResponse, 0, len(story.Pages))
	for _, page := range story.Pages {
		pages = append(pages, PageResponse{
			ID:           page.ID,
			Number:       page.PageNumber,
			ContentCN:    page.ContentCN,
			ContentHakka: page.ContentHakka,
		})
	}

	return FullStoryResponse{
		StoryResponse: toResponse(story),
		Pages:         pages,
	}
}
