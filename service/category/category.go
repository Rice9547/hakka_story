package scategory

import dcategory "github.com/rice9547/hakka_story/domain/category"

type (
	Service interface {
		Create(c *dcategory.Category) (*dcategory.Category, error)
		ListByName(name string) ([]dcategory.Category, error)
		Update(id uint64, name string) (*dcategory.Category, error)
		DeleteByID(id uint64) error
	}

	service struct {
		repo dcategory.Repository
	}
)

func New(repo dcategory.Repository) Service {
	return &service{repo: repo}
}

func (s *service) Create(c *dcategory.Category) (*dcategory.Category, error) {
	if err := s.repo.Save(c); err != nil {
		return nil, err
	}

	return c, nil
}

func (s *service) ListByName(name string) ([]dcategory.Category, error) {
	categories, err := s.repo.ListByKeyword(name)
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (s *service) Update(id uint64, name string) (*dcategory.Category, error) {
	category := &dcategory.Category{
		ID:   id,
		Name: name,
	}

	err := s.repo.UpdateByID(id, category)
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (s *service) DeleteByID(id uint64) error {
	return s.repo.DeleteByID(id)
}
