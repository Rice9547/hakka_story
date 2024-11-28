package sstory

import (
	dstory "github.com/rice9547/hakka_story/domain/story"
)

type Service interface {
	CreateStory(s *dstory.Story) error
	ListStory() ([]dstory.Story, error)
	ListStoryByCategories(categoryNames []string) ([]dstory.Story, error)
	GetStory(id uint64) (*dstory.Story, error)
	UpdateByID(id uint64, s *dstory.Story) error
	DeleteByID(id uint64) error
}

type service struct {
	repo dstory.Repository
}

func New(repo dstory.Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreateStory(st *dstory.Story) error {
	return s.repo.Save(st)
}

func (s *service) ListStory() ([]dstory.Story, error) {
	return s.repo.List()
}

func (s *service) ListStoryByCategories(categoryNames []string) ([]dstory.Story, error) {
	return s.repo.FilterByCategories(categoryNames)
}

func (s *service) GetStory(id uint64) (*dstory.Story, error) {
	return s.repo.GetByID(id)
}

func (s *service) UpdateByID(id uint64, st *dstory.Story) error {
	return s.repo.UpdateByID(id, st)
}

func (s *service) DeleteByID(id uint64) error {
	return s.repo.DeleteByID(id)
}
