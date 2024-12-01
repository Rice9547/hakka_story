package scategory

import (
	"context"

	dcategory "github.com/rice9547/hakka_story/domain/category"
	"github.com/rice9547/hakka_story/entities"
)

type (
	Service interface {
		Create(ctx context.Context, c *entities.Category) (*entities.Category, error)
		ListByName(ctx context.Context, name string) ([]entities.Category, error)
		Update(ctx context.Context, id uint64, name string) (*entities.Category, error)
		DeleteByID(ctx context.Context, id uint64) error
	}

	service struct {
		repo dcategory.Repository
	}
)

func New(repo dcategory.Repository) Service {
	return &service{repo: repo}
}

func (s *service) Create(ctx context.Context, c *entities.Category) (*entities.Category, error) {
	if err := s.repo.Save(ctx, c); err != nil {
		return nil, err
	}

	return c, nil
}

func (s *service) ListByName(ctx context.Context, name string) ([]entities.Category, error) {
	categories, err := s.repo.ListByKeyword(ctx, name)
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (s *service) Update(ctx context.Context, id uint64, name string) (*entities.Category, error) {
	category := &entities.Category{
		ID:   id,
		Name: name,
	}

	err := s.repo.UpdateByID(ctx, id, category)
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (s *service) DeleteByID(ctx context.Context, id uint64) error {
	return s.repo.DeleteByID(ctx, id)
}
