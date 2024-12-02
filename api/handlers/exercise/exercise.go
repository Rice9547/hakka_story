package hexercise

import (
	"context"

	"github.com/rice9547/hakka_story/entities"
	"github.com/rice9547/hakka_story/repository"
)

type (
	Service interface {
		CreateExercise(ctx context.Context, exercise *entities.Exercise) error
		GetExerciseCountByStoryIDs(ctx context.Context, storyIDs []uint64) ([]repository.ExerciseCounter, error)
		ListExerciseByStoryID(ctx context.Context, storyID uint64) ([]entities.Exercise, error)
		UpdateExercise(ctx context.Context, storyID, exerciseID uint64, exercise *entities.Exercise) error
	}

	Exercise struct {
		service Service
	}

	ChoiceResponse struct {
		ID         uint64 `json:"id"`
		ChoiceText string `json:"choice_text"`
		IsCorrect  *bool  `json:"is_correct,omitempty"`
	}

	OpenAnswerResponse struct {
		ID         uint64 `json:"id"`
		AnswerText string `json:"answer_text"`
	}

	ExerciseResponse struct {
		ID         uint64                `json:"id"`
		StoryID    uint64                `json:"story_id"`
		Type       entities.ExerciseType `json:"type"`
		PromptText string                `json:"prompt_text"`
		AudioURL   string                `json:"audio_url"`
		Choices    []ChoiceResponse      `json:"choices"`
		Answers    []OpenAnswerResponse  `json:"answers,omitempty"`
	}
)

func New(service Service) *Exercise {
	return &Exercise{service: service}
}

func toExerciseResponse(exercise entities.Exercise, isAdmin bool) ExerciseResponse {
	choices := make([]ChoiceResponse, 0, len(exercise.Choices))
	for _, choice := range exercise.Choices {
		choices = append(choices, ChoiceResponse{
			ID:         choice.ID,
			ChoiceText: choice.ChoiceText,
		})

		if isAdmin {
			choices[len(choices)-1].IsCorrect = &choice.IsCorrect
		}
	}

	answers := make([]OpenAnswerResponse, 0, len(exercise.Answers))
	for _, answer := range exercise.Answers {
		answers = append(answers, OpenAnswerResponse{
			ID:         answer.ID,
			AnswerText: answer.AnswerText,
		})
	}

	resp := ExerciseResponse{
		ID:         exercise.ID,
		StoryID:    exercise.StoryID,
		Type:       exercise.Type,
		PromptText: exercise.PromptText,
		AudioURL:   exercise.AudioURL,
		Choices:    choices,
	}

	if isAdmin {
		resp.Answers = answers
	}

	return resp
}
