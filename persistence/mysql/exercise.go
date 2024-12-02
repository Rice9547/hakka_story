package mysql

import (
	"context"
	"github.com/rice9547/hakka_story/entities"
	"github.com/rice9547/hakka_story/repository"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ExerciseRepository struct {
	DB *gorm.DB
}

func NewExercise(client *Client) repository.Exercise {
	return &ExerciseRepository{DB: client.DB()}
}

func (r *ExerciseRepository) CountMany(ctx context.Context, storyIDs []uint64) ([]repository.ExerciseCounter, error) {
	result := make([]repository.ExerciseCounter, 0)

	query := r.DB.WithContext(ctx).
		Model(&entities.Exercise{}).
		Select("exercises.story_id as story_id," +
			"stories.title as story_title," +
			"COUNT(*) as count").
		Joins("INNER JOIN stories ON stories.id = exercises.story_id")

	if len(storyIDs) > 0 {
		query = query.Where("story_id IN (?)", storyIDs)
	}

	err := query.Group("story_id").Find(&result).Error

	return result, err
}

func (r *ExerciseRepository) List(ctx context.Context, storyID uint64) ([]entities.Exercise, error) {
	exercises := make([]entities.Exercise, 0)
	err := r.DB.WithContext(ctx).
		Preload(clause.Associations).
		Where("story_id = ?", storyID).
		Find(&exercises).Error
	return exercises, err
}
