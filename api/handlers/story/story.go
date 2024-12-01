package hstory

import (
	"context"

	"github.com/rice9547/hakka_story/entities"
)

type (
	Service interface {
		CreateStory(ctx context.Context, s *entities.Story) error
		ListStory(ctx context.Context) ([]entities.Story, error)
		ListStoryByCategories(ctx context.Context, categoryNames []string) ([]entities.Story, error)
		GetStory(ctx context.Context, id uint64) (*entities.Story, error)
		UpdateByID(ctx context.Context, id uint64, s *entities.Story) error
		DeleteByID(ctx context.Context, id uint64) error
	}

	Story struct {
		service Service
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
		Image        string          `json:"image"`
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

func New(service Service) *Story {
	return &Story{service: service}
}

func toResponse(story entities.Story) StoryResponse {
	resp := StoryResponse{
		ID:          story.ID,
		Title:       story.Title,
		Description: story.Description,
		CoverImage:  story.Image,
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

func toFullyResponse(story entities.Story) FullStoryResponse {
	pages := make([]PageResponse, 0, len(story.Pages))
	for _, page := range story.Pages {
		currentPage := PageResponse{
			ID:           page.ID,
			Number:       page.PageNumber,
			Image:        page.Image,
			ContentCN:    page.ContentCN,
			ContentHakka: page.ContentHakka,
			Audios:       make([]AudioResponse, 0, len(page.AudioFiles)),
		}

		for _, audio := range page.AudioFiles {
			currentPage.Audios = append(currentPage.Audios, AudioResponse{
				ID:       audio.ID,
				AudioURL: audio.AudioURL,
				Dialect:  audio.Dialect,
			})
		}

		pages = append(pages, currentPage)
	}

	return FullStoryResponse{
		StoryResponse: toResponse(story),
		Pages:         pages,
	}
}
