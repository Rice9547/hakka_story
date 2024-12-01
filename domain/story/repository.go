package dstory

import (
	"context"
	"github.com/rice9547/hakka_story/entities"
)

type Repository interface {
	Save(ctx context.Context, s *entities.Story) error
	List(ctx context.Context) ([]entities.Story, error)
	FilterByCategories(ctx context.Context, categoryNames []string) ([]entities.Story, error)
	GetByID(ctx context.Context, id uint64) (*entities.Story, error)
	UpdateByID(ctx context.Context, id uint64, s *entities.Story) error
	DeleteByID(ctx context.Context, id uint64) error
}
