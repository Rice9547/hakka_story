package dstory

import "context"

type Repository interface {
	Save(ctx context.Context, s *Story) error
	List(ctx context.Context) ([]Story, error)
	FilterByCategories(ctx context.Context, categoryNames []string) ([]Story, error)
	GetByID(ctx context.Context, id uint64) (*Story, error)
	UpdateByID(ctx context.Context, id uint64, s *Story) error
	DeleteByID(ctx context.Context, id uint64) error
}
