package dstory

type Repository interface {
	Save(s *Story) error
	List() ([]Story, error)
	FilterByCategories(categoryNames []string) ([]Story, error)
	GetByID(id uint64) (*Story, error)
	UpdateByID(id uint64, s *Story) error
	DeleteByID(id uint64) error
}
