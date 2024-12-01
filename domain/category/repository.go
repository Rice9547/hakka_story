package dcategory

import "context"

type Repository interface {
	Save(ctx context.Context, c *Category) error
	ListByKeyword(ctx context.Context, keyword string) ([]Category, error)
	UpdateByID(ctx context.Context, id uint64, s *Category) error
	DeleteByID(ctx context.Context, id uint64) error
}
