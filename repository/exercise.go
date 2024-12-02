package repository

import (
	"context"
	"github.com/rice9547/hakka_story/entities"
)

type (
	Exercise interface {
		Save(ctx context.Context, exercise *entities.Exercise) error
		CountMany(ctx context.Context, storyIDs []uint64) ([]ExerciseCounter, error)
		List(ctx context.Context, storyID uint64) ([]entities.Exercise, error)
		Update(ctx context.Context, exerciseID uint64, exercise *entities.Exercise) error
	}

	ExerciseCounter struct {
		StoryID    uint64
		StoryTitle string
		Count      int64
	}
)