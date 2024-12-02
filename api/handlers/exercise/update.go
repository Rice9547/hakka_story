package hexercise

import (
	"github.com/gin-gonic/gin"
	"github.com/rice9547/hakka_story/lib/response"
	"net/http"
)

// UpdateExercise godoc
// @Summary Update an exercise
// @Description Update an exists exercise associated with a specific story ID
// @Tags admin exercise
// @Accept json
// @Produce json
// @Param id path uint64 true "Story ID"
// @Param exercise_id path uint64 true "Exercise ID"
// @Param exercise body UpsertExerciseRequest true "Exercise data"
// @Success 200 {object} response.Response{data=ExerciseAdminResponse}
// @Failure 400 {object} response.ResponseBase "Invalid input"
// @Failure 500 {object} response.ResponseBase "Internal server error"
// @Router /admin/story/{id}/exercise/{exercise_id} [put]
func (h *Exercise) UpdateExercise(c *gin.Context) {
	req := new(UpsertExerciseRequest)
	exercise, err := req.bind(c)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid input")
		return
	}

	if err := h.service.UpdateExercise(c, exercise.StoryID, exercise.ID, exercise); err != nil {
		response.Error(c, http.StatusInternalServerError, "Internal server error")
		return
	}

	response.Success(c, toExerciseAdminResponse(*exercise))
}
