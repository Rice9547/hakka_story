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
		DeleteExercise(ctx context.Context, storyID, exerciseID uint64) error
	}

	Exercise struct {
		service Service
	}

	ChoiceBaseResponse struct {
		ID         uint64 `json:"id"`
		ChoiceText string `json:"choice_text"`
	}

	ChoiceResponse struct {
		ChoiceBaseResponse
		IsCorrect bool `json:"is_correct"`
	}

	OpenAnswerResponse struct {
		ID         uint64 `json:"id"`
		AnswerText string `json:"answer_text"`
	}

	ExerciseBaseResponse struct {
		ID         uint64                `json:"id"`
		StoryID    uint64                `json:"story_id"`
		Type       entities.ExerciseType `json:"type"`
		PromptText string                `json:"prompt_text"`
		AudioURL   string                `json:"audio_url"`
	}

	ExerciseAdminResponse struct {
		ExerciseBaseResponse
		Choices []ChoiceResponse     `json:"choices"`
		Answers []OpenAnswerResponse `json:"answers,omitempty"`
	}

	ExerciseResponse struct {
		ExerciseBaseResponse
		Choices []ChoiceBaseResponse `json:"choices"`
	}
)

func New(service Service) *Exercise {
	return &Exercise{service: service}
}

func toExerciseBaseResponse(exercise entities.Exercise) ExerciseBaseResponse {
	choices := make([]ChoiceBaseResponse, 0, len(exercise.Choices))
	for _, choice := range exercise.Choices {
		choices = append(choices, ChoiceBaseResponse{
			ID:         choice.ID,
			ChoiceText: choice.ChoiceText,
		})
	}

	return ExerciseBaseResponse{
		ID:         exercise.ID,
		StoryID:    exercise.StoryID,
		Type:       exercise.Type,
		PromptText: exercise.PromptText,
		AudioURL:   exercise.AudioURL,
	}
}

func toExerciseResponse(exercise entities.Exercise) ExerciseResponse {
	base := toExerciseBaseResponse(exercise)

	choices := make([]ChoiceBaseResponse, 0, len(exercise.Choices))
	for _, choice := range exercise.Choices {
		choices = append(choices, ChoiceBaseResponse{
			ID:         choice.ID,
			ChoiceText: choice.ChoiceText,
		})
	}

	return ExerciseResponse{
		ExerciseBaseResponse: base,
		Choices:              choices,
	}
}

func toExerciseAdminResponse(exercise entities.Exercise) ExerciseAdminResponse {
	base := toExerciseBaseResponse(exercise)

	choices := make([]ChoiceResponse, 0, len(exercise.Choices))
	for _, choice := range exercise.Choices {
		choices = append(choices, ChoiceResponse{
			ChoiceBaseResponse: ChoiceBaseResponse{
				ID:         choice.ID,
				ChoiceText: choice.ChoiceText,
			},
			IsCorrect: choice.IsCorrect,
		})
	}

	answers := make([]OpenAnswerResponse, 0, len(exercise.Answers))
	for _, answer := range exercise.Answers {
		answers = append(answers, OpenAnswerResponse{
			ID:         answer.ID,
			AnswerText: answer.AnswerText,
		})
	}

	return ExerciseAdminResponse{
		ExerciseBaseResponse: base,
		Choices:              choices,
		Answers:              answers,
	}
}
