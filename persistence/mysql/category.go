package mysql

import (
	"gorm.io/gorm"

	dcategory "github.com/rice9547/hakka_story/domain/category"
	"github.com/rice9547/hakka_story/lib/errors"
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
	err := r.DB.Model(&dcategory.Category{}).
		Where("name LIKE ?", "%"+keyword+"%").
		Find(&categories).Error
	return categories, err
}

func (r *CategoryRepository) UpdateByID(id uint64, category *dcategory.Category) error {
	result := r.DB.Model(&dcategory.Category{}).
		Where("id = ?", id).
		Update("name", category.Name)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.ErrCategoryNotFound
	}

	return nil
}

func (r *CategoryRepository) DeleteByID(id uint64) error {
	result := r.DB.Delete(&dcategory.Category{}, id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.ErrCategoryNotFound
	}

	return nil
}
