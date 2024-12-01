package dcategory

import (
	"context"

	"github.com/rice9547/hakka_story/entities"
)

type Repository interface {
	Save(ctx context.Context, c *entities.Category) error
	ListByKeyword(ctx context.Context, keyword string) ([]entities.Category, error)
	UpdateByID(ctx context.Context, id uint64, s *entities.Category) error
	DeleteByID(ctx context.Context, id uint64) error
}
