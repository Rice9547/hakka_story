package hexercise

import (
	"github.com/rice9547/hakka_story/lib/errors"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/rice9547/hakka_story/entities"
	"github.com/rice9547/hakka_story/lib/response"
)

type (
	DoExerciseRequest struct {
		id     uint64
		userID string

		Type    entities.ExerciseType `json:"type" binding:"gte=0"`
		Answers []string              `json:"answers" binding:"required"`
	}

	DoExerciseResponse struct {
		IsCorrect bool     `json:"is_correct"`
		Answers   []string `json:"answers"`
	}
)

func (r *DoExerciseRequest) bind(c *gin.Context) (*DoExerciseRequest, error) {
	if err := c.ShouldBindJSON(r); err != nil {
		return nil, err
	}

	exerciseID, err := strconv.ParseUint(c.Param("exercise_id"), 10, 64)
	if err != nil {
		return nil, err
	}

	r.id = exerciseID

	if userID, exists := c.Get("user"); exists {
		r.userID = userID.(string)
	}

	return r, nil
}

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
		response.BadRequest(c, err, "Invalid input")
		return
	}

	if err := h.service.UpdateExercise(c, exercise.StoryID, exercise.ID, exercise); err != nil {
		response.InternalServerError(c, err, "Failed to update exercise")
		return
	}

	response.Success(c, toExerciseAdminResponse(*exercise))
}

// DoExercise godoc
// @Summary Do an exercise
// @Description Do an exists exercise associated with a specific story ID
// @Tags exercise
// @Accept json
// @Produce json
// @Param exercise_id path uint64 true "Exercise ID"
// @Param exercise body DoExerciseRequest true "Exercise data"
// @Success 200 {object} DoExerciseResponse
// @Failure 400 {object} response.ResponseBase "Invalid input"
// @Failure 400 {object} response.ResponseBase "Exercise type mismatch"
// @Failure 404 {object} response.ResponseBase "Exercise not found"
// @Failure 500 {object} response.ResponseBase "Internal server error"
// @Router /exercise/{exercise_id} [post]
func (h *Exercise) Do(c *gin.Context) {
	req := new(DoExerciseRequest)
	exercise, err := req.bind(c)
	if err != nil {
		response.BadRequest(c, err, "Invalid input")
		return
	}

	isCorrect, correctAnswers, err := h.service.DoExercise(c, exercise.userID, exercise.id, exercise.Type, exercise.Answers)
	if err != nil {
		switch true {
		case errors.Is(err, errors.ErrExerciseTypeMismatch):
			response.BadRequest(c, err, "Exercise type mismatch")
			return
		case errors.Is(err, errors.ErrExerciseNotFound):
			response.NotFound(c, "Exercise not found")
			return
		case errors.Is(err, errors.ErrExerciseTypeNotSupport):
			response.InternalServerError(c, err, "Exercise type not support")
			return
		}
		response.InternalServerError(c, err, "Failed to do exercise")
		return
	}

	response.Success(c, DoExerciseResponse{
		IsCorrect: isCorrect,
		Answers:   correctAnswers,
	})
}
