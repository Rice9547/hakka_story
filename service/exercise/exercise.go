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
