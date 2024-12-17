package sexercise

import (
	"context"
	"github.com/rice9547/hakka_story/entities"
	"github.com/rice9547/hakka_story/lib/errors"
	"github.com/rice9547/hakka_story/repository"
	"slices"
)

type Exercise struct {
	exerciseRepo repository.Exercise
}

func New(exerciseRepo repository.Exercise) *Exercise {
	return &Exercise{exerciseRepo: exerciseRepo}
}

func (e *Exercise) CreateExercise(ctx context.Context, exercise *entities.Exercise) error {
	return e.exerciseRepo.Save(ctx, exercise)
}

func (e *Exercise) GetExerciseCountByStoryIDs(ctx context.Context, storyIDs []uint64) ([]repository.ExerciseCounter, error) {
	return e.exerciseRepo.CountMany(ctx, storyIDs)
}

func (e *Exercise) ListExerciseByStoryID(ctx context.Context, storyID uint64) ([]entities.Exercise, error) {
	return e.exerciseRepo.List(ctx, storyID)
}

func (e *Exercise) ListExerciseByStoryIDs(ctx context.Context, storyIDs []uint64) ([]entities.Exercise, error) {
	return e.exerciseRepo.ListMany(ctx, storyIDs)
}

func (e *Exercise) UpdateExercise(ctx context.Context, storyID, exerciseID uint64, exercise *entities.Exercise) error {
	exercise.StoryID = storyID
	return e.exerciseRepo.Update(ctx, exerciseID, exercise)
}

func (e *Exercise) DeleteExercise(ctx context.Context, storyID, exerciseID uint64) error {
	return e.exerciseRepo.Delete(ctx, storyID, exerciseID)
}

func (e *Exercise) DoExercise(
	ctx context.Context, userID string, exerciseID uint64, typ entities.ExerciseType, answers []string,
) (bool, []string, error) {
	exercise, err := e.exerciseRepo.Get(ctx, exerciseID)
	if err != nil {
		return false, nil, err
	}

	if exercise.Type != typ {
		return false, nil, errors.ErrExerciseTypeMismatch
	}

	var (
		isCorrect      bool
		correctAnswers []string
	)
	switch typ {
	case entities.ExerciseTypeChoice:
		isCorrect, correctAnswers, err = e.doChoiceExercise(exercise, answers)
	case entities.ExerciseTypeFillIn:
		isCorrect, correctAnswers, err = e.doFillInExercise(exercise, answers)
	default:
		return false, nil, errors.ErrExerciseTypeNotSupport
	}

	if userID != "" {
		// TODO:Save user's answer
	}

	return isCorrect, correctAnswers, err
}

func (e *Exercise) doChoiceExercise(exercise *entities.Exercise, answers []string) (bool, []string, error) {
	if len(answers) == 0 || len(answers) > len(exercise.Choices) {
		return false, nil, errors.ErrExerciseInvalidAnswer
	}

	correctAnswers := make([]string, 0, len(exercise.Choices))
	for _, choice := range exercise.Choices {
		if choice.IsCorrect {
			correctAnswers = append(correctAnswers, choice.ChoiceText)
		}
	}

	slices.Sort(correctAnswers)
	slices.Sort(answers)
	isCorrect := slices.Equal(correctAnswers, answers)

	return isCorrect, correctAnswers, nil
}

func (e *Exercise) doFillInExercise(exercise *entities.Exercise, answers []string) (bool, []string, error) {
	if len(answers) != 1 {
		return false, nil, errors.ErrExerciseInvalidAnswer
	}

	correctAnswers := make([]string, 0, len(exercise.Answers))
	isCorrect := false
	for _, answer := range exercise.Answers {
		correctAnswers = append(correctAnswers, answer.AnswerText)
		if answer.AnswerText == answers[0] {
			isCorrect = true
		}
	}

	return isCorrect, correctAnswers, nil
}
