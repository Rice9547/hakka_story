package hexercise

import (
	"github.com/gin-gonic/gin"
	"github.com/rice9547/hakka_story/entities"
	"github.com/rice9547/hakka_story/lib/response"
	"net/http"
	"strconv"
)

type (
	UpsertChoiceRequest struct {
		ChoiceText string `json:"choice_text" binding:"required"`
		IsCorrect  bool   `json:"is_correct" binding:"required"`
	}

	UpsertAnswerRequest struct {
		AnswerText string `json:"answer_text" binding:"required"`
	}

	UpsertExerciseRequest struct {
		PromptText string                `json:"prompt_text" binding:"required"`
		Type       int                   `json:"type"`
		AudioURL   string                `json:"audio_url"`
		Choices    []UpsertChoiceRequest `json:"choices"`
		Answers    []UpsertAnswerRequest `json:"answers"`
	}
)

func (r *UpsertExerciseRequest) bind(c *gin.Context) (*entities.Exercise, error) {
	if err := c.ShouldBind(r); err != nil {
		return nil, err
	}

	storyID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return nil, err
	}

	exercise := &entities.Exercise{
		StoryID:    storyID,
		PromptText: r.PromptText,
		Type:       entities.ExerciseType(r.Type),
		AudioURL:   r.AudioURL,
		Choices:    make([]entities.ExerciseChoice, 0, len(r.Choices)),
		Answers:    make([]entities.ExerciseOpenAnswer, 0, len(r.Answers)),
	}

	if c.Param("exercise_id") != "" {
		exerciseID, err := strconv.ParseUint(c.Param("exercise_id"), 10, 64)
		if err != nil {
			return nil, err
		}
		exercise.ID = exerciseID
	}

	for _, choice := range r.Choices {
		exercise.Choices = append(exercise.Choices, entities.ExerciseChoice{
			ChoiceText: choice.ChoiceText,
			IsCorrect:  choice.IsCorrect,
		})
	}

	for _, answer := range r.Answers {
		exercise.Answers = append(exercise.Answers, entities.ExerciseOpenAnswer{
			AnswerText: answer.AnswerText,
		})
	}

	return exercise, nil
}

// CreateExercise godoc
// @Summary Create a new exercise
// @Description Creates a new exercise associated with a specific story ID
// @Tags admin exercise
// @Accept json
// @Produce json
// @Param id path uint64 true "Story ID"
// @Param exercise body UpsertExerciseRequest true "Exercise data"
// @Success 200 {object} response.Response{data=ExerciseResponse}
// @Failure 400 {object} response.ResponseBase "Invalid input"
// @Failure 500 {object} response.ResponseBase "Internal server error"
// @Router /admin/story/{id}/exercise [post]
func (h *Exercise) CreateExercise(c *gin.Context) {
	req := new(UpsertExerciseRequest)
	exercise, err := req.bind(c)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid input")
		return
	}

	if err := h.service.CreateExercise(c, exercise); err != nil {
		response.Error(c, http.StatusInternalServerError, "Internal server error")
		return
	}

	response.Success(c, toExerciseResponse(*exercise, true))
}
