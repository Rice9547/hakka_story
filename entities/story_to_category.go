package entities

import "time"

type StoryToCategory struct {
	StoryID    uint64 `gorm:"column:story_id"`
	CategoryID uint64 `gorm:"column:category_id"`
	CreatedAt  time.Time
}

func (StoryToCategory) TableName() string {
	return "story_to_category"
}
