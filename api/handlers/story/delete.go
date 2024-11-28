package hstory

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/rice9547/hakka_story/lib/errors"
	"github.com/rice9547/hakka_story/lib/response"
)

// DeleteStory godoc
// @Summary Delete a story by ID
// @Description Delete a story by its ID
// @Tags admin story
// @Accept json
// @Produce json
// @Param id path uint64 true "Story ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.ResponseBase
// @Failure 404 {object} response.ResponseBase
// @Failure 500 {object} response.ResponseBase
// @Router /admin/story/{id} [delete]
func (h *Story) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		errors.ErrorHandler(c, errors.NewAppError(http.StatusBadRequest, errors.ErrInvalidInput, "Invalid story ID"))
		return
	}

	if err = h.service.DeleteByID(id); err != nil {
		if errors.Is(err, errors.ErrStoryNotFound) {
			errors.ErrorHandler(c, errors.NewAppError(http.StatusNotFound, err, "Story not found"))
			return
		}
		errors.ErrorHandler(c, errors.NewAppError(http.StatusInternalServerError, err, "Failed to delete story"))
		return
	}

	response.Success(c, nil)
}
