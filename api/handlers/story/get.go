package hstory

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/rice9547/hakka_story/lib/errors"
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
		errors.ErrorHandler(c, errors.NewAppError(http.StatusBadRequest, errors.ErrInvalidInput, "Invalid story ID"))
		return
	}

	story, err := h.service.GetStory(c.Request.Context(), id)
	if err != nil {
		errors.ErrorHandler(c, errors.NewAppError(http.StatusNotFound, err, "Story not found"))
		return
	}

	response.Success(c, toFullyResponse(*story))
}
