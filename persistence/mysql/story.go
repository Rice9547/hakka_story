package mysql

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	dstory "github.com/rice9547/hakka_story/domain/story"
	"github.com/rice9547/hakka_story/lib/errors"
)

type StoryRepository struct {
	DB *gorm.DB
}

func NewStory(client *Client) dstory.Repository {
	return &StoryRepository{DB: client.DB()}
}

func (r *StoryRepository) Save(s *dstory.Story) error {
	return r.DB.Save(s).Error
}

func (r *StoryRepository) List() ([]dstory.Story, error) {
	stories := make([]dstory.Story, 0)
	err := r.DB.Preload(clause.Associations).Find(&stories).Error
	return stories, err
}

func (r *StoryRepository) GetByID(id uint64) (*dstory.Story, error) {
	story := &dstory.Story{}
	err := r.DB.Model(&story).Preload(clause.Associations).First(&story, id).Error
	return story, err
}

func (r *StoryRepository) UpdateByID(id uint64, story *dstory.Story) error {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		ori := new(dstory.Story)
		result := tx.Model(&dstory.Story{}).Where("id = ?", id).First(ori)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				return errors.ErrStoryNotFound
			}
			return result.Error
		}

		if err := tx.Model(&dstory.StoryPage{}).Delete(&dstory.StoryPage{}, "story_id = ?", id).Error; err != nil {
			return err
		}

		story.ID = id
		story.CreatedAt = ori.CreatedAt

		return tx.Save(story).Error
	})
}
