package hstory

import (
	"github.com/gin-gonic/gin"
	"github.com/rice9547/hakka_story/entities"

	"github.com/rice9547/hakka_story/lib/response"
)

type (
	ListStoryResponse struct {
		Stories []StoryResponse `json:"stories"`
	}
)

// GetStories godoc
// @Summary      List stories
// @Description  Get all stories
// @Tags         story
// @Produce      json
// @Param        name  query    []string  false  "Category names"
// @Success      200  {object}  response.Response{data=StoryResponse}
// @Failure      500  {object}  response.ResponseBase
// @Router       /story [get]
func (h *Story) List(c *gin.Context) {
	categoryNames := c.QueryArray("categories")

	var (
		stories []entities.Story
		err     error
	)
	if len(categoryNames) == 0 {
		stories, err = h.service.ListStory(c.Request.Context())
	} else {
		stories, err = h.service.ListStoryByCategories(c.Request.Context(), categoryNames)
	}

	if err != nil {
		response.InternalServerError(c, err, "Failed to retrieve stories")
		return
	}

	resp := make([]StoryResponse, 0, len(stories))
	for _, story := range stories {
		resp = append(resp, toResponse(story))
	}

	response.Success(c, resp)
}
