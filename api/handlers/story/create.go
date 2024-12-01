package hstory

import (
	"net/http"

	"github.com/gin-gonic/gin"

	dcategory "github.com/rice9547/hakka_story/domain/category"
	dstory "github.com/rice9547/hakka_story/domain/story"
	"github.com/rice9547/hakka_story/lib/errors"
	"github.com/rice9547/hakka_story/lib/response"
)

type (
	UpsertAudioRequest struct {
		AudioURL string `json:"audio_url" binding:"required"`
		Dialect  string `json:"dialect" binding:"required"`
	}

	UpsertPageRequest struct {
		Number       int                  `json:"page_number" binding:"required"`
		ContentCN    string               `json:"content_cn" binding:"required"`
		ContentHakka string               `json:"content_hakka" binding:"required"`
		Audios       []UpsertAudioRequest `json:"audios"`
	}

	UpsertCategoryRequest struct {
		ID   uint64 `json:"id"`
		Name string `json:"name"`
	}

	UpsertStoryRequest struct {
		Title       string                  `json:"title" binding:"required"`
		Description string                  `json:"description" binding:"required"`
		Pages       []UpsertPageRequest     `json:"pages" binding:"required"`
		CoverImage  string                  `json:"cover_image"`
		Categories  []UpsertCategoryRequest `json:"categories"`
	}
)

func (s *UpsertStoryRequest) bind(ctx *gin.Context) (*dstory.Story, error) {
	if err := ctx.ShouldBind(s); err != nil {
		return nil, err
	}

	story := &dstory.Story{
		Title:       s.Title,
		Description: s.Description,
		Image:       &dstory.Image{ImageURL: s.CoverImage},
		Pages:       make([]dstory.StoryPage, 0, len(s.Pages)),
		Categories:  make([]dcategory.Category, 0, len(s.Categories)),
	}

	for _, page := range s.Pages {
		story.Pages = append(story.Pages, dstory.StoryPage{
			PageNumber:   page.Number,
			ContentCN:    page.ContentCN,
			ContentHakka: page.ContentHakka,
			AudioFiles:   make([]dstory.AudioFile, 0, len(page.Audios)),
		})

		for _, audio := range page.Audios {
			story.Pages[len(story.Pages)-1].AudioFiles = append(story.Pages[len(story.Pages)-1].AudioFiles, dstory.AudioFile{
				AudioURL: audio.AudioURL,
				Dialect:  audio.Dialect,
			})
		}
	}

	for _, category := range s.Categories {
		story.Categories = append(story.Categories, dcategory.Category{
			ID:   category.ID,
			Name: category.Name,
		})
	}

	return story, nil
}

// CreateStory godoc
// @Summary      Create Story
// @Description  Create Story
// @Tags         admin story
// @Accept       json
// @Produce      json
// @Param		 Authorization  header  string  true  "Bearer token"
// @Param        story  body      UpsertStoryRequest  true  "故事內容"
// @Success      201    {object}  response.Response{data=StoryResponse}
// @Failure      400    {object}  response.ResponseBase
// @Failure      401    {object}  response.ResponseBase
// @Failure      500    {object}  response.ResponseBase
// @Router       /admin/story [post]
func (h *Story) Create(c *gin.Context) {
	request := new(UpsertStoryRequest)
	story, err := request.bind(c)
	if err != nil {
		errors.ErrorHandler(c, errors.NewAppError(http.StatusBadRequest, errors.ErrInvalidInput, "Invalid input"))
		return
	}

	if err = h.service.CreateStory(c.Request.Context(), story); err != nil {
		errors.ErrorHandler(c, errors.NewAppError(http.StatusInternalServerError, err, "Failed to create story"))
		return
	}

	response.Success(c, toResponse(*story))
}
