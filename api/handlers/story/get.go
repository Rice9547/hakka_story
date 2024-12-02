package hstory

import (
	"github.com/rice9547/hakka_story/lib/errors"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/rice9547/hakka_story/lib/response"
)

// GetStory godoc
// @Summary      Get story
// @Description  Get story by id with pages
// @Tags         stories
// @Produce      json
// @Param        id   path      int  true  "Story ID"
// @Success      200  {object}  response.Response{data=FullStoryResponse}
// @Failure      500  {object}  response.ResponseBase
// @Router       /story/:id [get]
func (h *Story) Get(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.BadRequest(c, err, "Invalid story ID")
		return
	}

	story, err := h.service.GetStory(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, errors.ErrStoryNotFound) {
			response.NotFound(c, "Story not found")
			return
		}
		response.InternalServerError(c, err, "Failed to get story")
		return
	}

	response.Success(c, toFullyResponse(*story))
}
