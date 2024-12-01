package sstory

import (
	"context"

	dstory "github.com/rice9547/hakka_story/domain/story"
	"github.com/rice9547/hakka_story/entities"
)

type Service interface {
	CreateStory(ctx context.Context, s *entities.Story) error
	ListStory(ctx context.Context) ([]entities.Story, error)
	ListStoryByCategories(ctx context.Context, categoryNames []string) ([]entities.Story, error)
	GetStory(ctx context.Context, id uint64) (*entities.Story, error)
	UpdateByID(ctx context.Context, id uint64, s *entities.Story) error
	DeleteByID(ctx context.Context, id uint64) error
}

type service struct {
	repo dstory.Repository
}

func New(repo dstory.Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreateStory(ctx context.Context, st *entities.Story) error {
	return s.repo.Save(ctx, st)
}

func (s *service) ListStory(ctx context.Context) ([]entities.Story, error) {
	return s.repo.List(ctx)
}

func (s *service) ListStoryByCategories(ctx context.Context, categoryNames []string) ([]entities.Story, error) {
	return s.repo.FilterByCategories(ctx, categoryNames)
}

func (s *service) GetStory(ctx context.Context, id uint64) (*entities.Story, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *service) UpdateByID(ctx context.Context, id uint64, st *entities.Story) error {
	return s.repo.UpdateByID(ctx, id, st)
}

func (s *service) DeleteByID(ctx context.Context, id uint64) error {
	return s.repo.DeleteByID(ctx, id)
}
