package sstory

import (
	"context"
	"github.com/rice9547/hakka_story/repository"

	"github.com/rice9547/hakka_story/entities"
)

type Service struct {
	storyRepo repository.Story
}

func New(storyRepo repository.Story) *Service {
	return &Service{storyRepo: storyRepo}
}

func (s *Service) CreateStory(ctx context.Context, st *entities.Story) error {
	return s.storyRepo.Save(ctx, st)
}

func (s *Service) ListStory(ctx context.Context) ([]entities.Story, error) {
	return s.storyRepo.List(ctx)
}

func (s *Service) ListStoryByCategories(ctx context.Context, categoryNames []string) ([]entities.Story, error) {
	return s.storyRepo.FilterByCategories(ctx, categoryNames)
}

func (s *Service) GetStory(ctx context.Context, id uint64) (*entities.Story, error) {
	return s.storyRepo.GetByID(ctx, id)
}

func (s *Service) UpdateByID(ctx context.Context, id uint64, st *entities.Story) error {
	return s.storyRepo.UpdateByID(ctx, id, st)
}

func (s *Service) DeleteByID(ctx context.Context, id uint64) error {
	return s.storyRepo.DeleteByID(ctx, id)
}
