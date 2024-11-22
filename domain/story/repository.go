package dstory

type Repository interface {
	Save(s *Story) error
	List() ([]Story, error)
	GetByID(id uint64) (*Story, error)
	UpdateByID(id uint64, s *Story) error
}
