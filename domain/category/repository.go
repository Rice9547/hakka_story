package dcategory

type Repository interface {
	Save(c *Category) error
	ListByKeyword(keyword string) ([]Category, error)
	UpdateByID(id uint64, s *Category) error
	DeleteByID(id uint64) error
}
