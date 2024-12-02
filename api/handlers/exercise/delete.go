package hexercise

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/rice9547/hakka_story/lib/errors"
	"github.com/rice9547/hakka_story/lib/response"
)

// DeleteExercise godoc
// @Summary DeleteExercise a exercise by ID
// @Description DeleteExercise an exercise by its ID
// @Tags admin exercise
// @Accept json
// @Produce json
// @Param id path uint64 true "Story ID"
// @Param exercise_id path uint64 true "Exercise ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.ResponseBase
// @Failure 404 {object} response.ResponseBase
// @Failure 500 {object} response.ResponseBase
// @Router /admin/story/{id}/exercise/{exercise_id} [delete]
func (h *Exercise) DeleteExercise(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		errors.ErrorHandler(c, errors.NewAppError(http.StatusBadRequest, errors.ErrInvalidInput, "Invalid story ID"))
		return
	}

	exerciseID, err := strconv.ParseUint(c.Param("exercise_id"), 10, 64)
	if err != nil {
		errors.ErrorHandler(c, errors.NewAppError(http.StatusBadRequest, errors.ErrInvalidInput, "Invalid exercise ID"))
		return
	}

	if err = h.service.DeleteExercise(c.Request.Context(), id, exerciseID); err != nil {
		if errors.Is(err, errors.ErrExerciseNotFound) {
			errors.ErrorHandler(c, errors.NewAppError(http.StatusNotFound, err, "Exercise not found"))
			return
		}
		errors.ErrorHandler(c, errors.NewAppError(http.StatusInternalServerError, err, "Failed to delete exercise"))
		return
	}

	response.Success(c, nil)
}
