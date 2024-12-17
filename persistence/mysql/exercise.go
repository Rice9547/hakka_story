package mysql

import (
	"context"
	"github.com/rice9547/hakka_story/entities"
	"github.com/rice9547/hakka_story/lib/errors"
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

func (r *ExerciseRepository) Save(ctx context.Context, exercise *entities.Exercise) error {
	return r.DB.WithContext(ctx).Save(exercise).Error
}

func (r *ExerciseRepository) Get(ctx context.Context, exerciseID uint64) (*entities.Exercise, error) {
	exercise := new(entities.Exercise)
	err := r.DB.WithContext(ctx).
		Preload(clause.Associations).
		First(exercise, exerciseID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.ErrExerciseNotFound
		}
		return nil, err
	}
	return exercise, nil
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

func (r *ExerciseRepository) ListMany(ctx context.Context, storyIDs []uint64) ([]entities.Exercise, error) {
	exercises := make([]entities.Exercise, 0)
	query := r.DB.WithContext(ctx).
		Preload(clause.Associations)

	if len(storyIDs) > 0 {
		query = query.Where("story_id IN (?)", storyIDs)
	}

	err := query.Find(&exercises).Error

	return exercises, err
}

func (r *ExerciseRepository) Update(ctx context.Context, exerciseID uint64, exercise *entities.Exercise) error {
	return r.DB.WithContext(ctx).
		Transaction(func(tx *gorm.DB) error {
			if err := tx.Delete(&entities.ExerciseChoice{}, "exercise_id = ?", exerciseID).Error; err != nil {
				return err
			}

			if err := tx.Delete(&entities.ExerciseOpenAnswer{}, "exercise_id = ?", exerciseID).Error; err != nil {
				return err
			}

			exercise.ID = exerciseID
			return tx.Save(exercise).Error
		})
}

func (r *ExerciseRepository) Delete(ctx context.Context, storyID, exerciseID uint64) error {
	result := r.DB.WithContext(ctx).
		Delete(&entities.Exercise{}, "story_id = ? AND id = ?", storyID, exerciseID)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.ErrExerciseNotFound
	}
	return nil
}
