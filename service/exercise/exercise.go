package sexercise

import (
	"context"
	"github.com/rice9547/hakka_story/entities"
	"github.com/rice9547/hakka_story/repository"
)

type Exercise struct {
	exerciseRepo repository.Exercise
}

func New(exerciseRepo repository.Exercise) *Exercise {
	return &Exercise{exerciseRepo: exerciseRepo}
}

func (e *Exercise) GetExerciseCountByStoryIDs(ctx context.Context, storyIDs []uint64) ([]repository.ExerciseCounter, error) {
	return e.exerciseRepo.CountMany(ctx, storyIDs)
}

func (e *Exercise) ListExerciseByStoryID(ctx context.Context, storyID uint64) ([]entities.Exercise, error) {
	return e.exerciseRepo.List(ctx, storyID)
}
