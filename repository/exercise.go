package repository

import (
	"context"
	"github.com/rice9547/hakka_story/entities"
)

type (
	Exercise interface {
		CountMany(ctx context.Context, storyIDs []uint64) ([]ExerciseCounter, error)
		List(ctx context.Context, storyID uint64) ([]entities.Exercise, error)
	}

	ExerciseCounter struct {
		StoryID    uint64
		StoryTitle string
		Count      int64
	}
)
