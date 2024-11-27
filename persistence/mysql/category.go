package mysql

import (
	"gorm.io/gorm"

	dcategory "github.com/rice9547/hakka_story/domain/category"
)

type CategoryRepository struct {
	DB *gorm.DB
}

func NewCategory(client *Client) dcategory.Repository {
	return &CategoryRepository{DB: client.DB()}
}

func (r *CategoryRepository) Save(c *dcategory.Category) error {
	return r.DB.Save(c).Error
}

func (r *CategoryRepository) ListByKeyword(keyword string) ([]dcategory.Category, error) {
	categories := make([]dcategory.Category, 0)
	// TODO: List categories by keyword
	return categories, nil
}

func (r *CategoryRepository) UpdateByID(id uint64, category *dcategory.Category) error {
	// TODO: Update category by ID
	return nil
}
