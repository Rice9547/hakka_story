package scategory

import (
	"context"
	"github.com/rice9547/hakka_story/repository"

	"github.com/rice9547/hakka_story/entities"
)

type (
	Service struct {
		categoryRepo repository.Category
	}
)

func New(categoryRepo repository.Category) *Service {
	return &Service{categoryRepo: categoryRepo}
}

func (s *Service) Create(ctx context.Context, c *entities.Category) (*entities.Category, error) {
	if err := s.categoryRepo.Save(ctx, c); err != nil {
		return nil, err
	}

	return c, nil
}

func (s *Service) ListByName(ctx context.Context, name string) ([]entities.Category, error) {
	categories, err := s.categoryRepo.ListByKeyword(ctx, name)
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (s *Service) Update(ctx context.Context, id uint64, name string) (*entities.Category, error) {
	category := &entities.Category{
		ID:   id,
		Name: name,
	}

	err := s.categoryRepo.UpdateByID(ctx, id, category)
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (s *Service) DeleteByID(ctx context.Context, id uint64) error {
	return s.categoryRepo.DeleteByID(ctx, id)
}
