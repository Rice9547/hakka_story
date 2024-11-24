package hstory

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/rice9547/hakka_story/lib/errors"
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
// @Tags         stories
// @Produce      json
// @Success      200  {object}  response.Response{data=StoryResponse}
// @Failure      500  {object}  response.ResponseBase
// @Router       /story [get]
func (h *Story) List(c *gin.Context) {
	stories, err := h.service.ListStory()
	if err != nil {
		errors.ErrorHandler(c, errors.NewAppError(http.StatusInternalServerError, err, "Failed to retrieve stories"))
		return
	}

	resp := make([]StoryResponse, 0, len(stories))
	for _, story := range stories {
		resp = append(resp, toResponse(story))
	}

	response.Success(c, resp)
}