package mysql

import (
	"context"
	"gorm.io/gorm"

	dcategory "github.com/rice9547/hakka_story/domain/category"
	"github.com/rice9547/hakka_story/entities"
	"github.com/rice9547/hakka_story/lib/errors"
)

type CategoryRepository struct {
	DB *gorm.DB
}

func NewCategory(client *Client) dcategory.Repository {
	return &CategoryRepository{DB: client.DB()}
}

func (r *CategoryRepository) Save(ctx context.Context, c *entities.Category) error {
	return r.DB.WithContext(ctx).Save(c).Error
}

func (r *CategoryRepository) ListByKeyword(ctx context.Context, keyword string) ([]entities.Category, error) {
	categories := make([]entities.Category, 0)
	err := r.DB.WithContext(ctx).
		Model(&entities.Category{}).
		Where("name LIKE ?", "%"+keyword+"%").
		Find(&categories).Error
	return categories, err
}

func (r *CategoryRepository) UpdateByID(ctx context.Context, id uint64, category *entities.Category) error {
	result := r.DB.WithContext(ctx).
		Model(&entities.Category{}).
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

func (r *CategoryRepository) DeleteByID(ctx context.Context, id uint64) error {
	result := r.DB.WithContext(ctx).
		Delete(&entities.Category{}, id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.ErrCategoryNotFound
	}

	return nil
}
