package sstory

import (
	"context"
	dstory "github.com/rice9547/hakka_story/domain/story"
)

type Service interface {
	CreateStory(ctx context.Context, s *dstory.Story) error
	ListStory(ctx context.Context) ([]dstory.Story, error)
	ListStoryByCategories(ctx context.Context, categoryNames []string) ([]dstory.Story, error)
	GetStory(ctx context.Context, id uint64) (*dstory.Story, error)
	UpdateByID(ctx context.Context, id uint64, s *dstory.Story) error
	DeleteByID(ctx context.Context, id uint64) error
}

type service struct {
	repo dstory.Repository
}

func New(repo dstory.Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreateStory(ctx context.Context, st *dstory.Story) error {
	return s.repo.Save(ctx, st)
}

func (s *service) ListStory(ctx context.Context) ([]dstory.Story, error) {
	return s.repo.List(ctx)
}

func (s *service) ListStoryByCategories(ctx context.Context, categoryNames []string) ([]dstory.Story, error) {
	return s.repo.FilterByCategories(ctx, categoryNames)
}

func (s *service) GetStory(ctx context.Context, id uint64) (*dstory.Story, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *service) UpdateByID(ctx context.Context, id uint64, st *dstory.Story) error {
	return s.repo.UpdateByID(ctx, id, st)
}

func (s *service) DeleteByID(ctx context.Context, id uint64) error {
	return s.repo.DeleteByID(ctx, id)
}
