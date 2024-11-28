package dstory

import (
	"time"

	"gorm.io/gorm"

	dcategory "github.com/rice9547/hakka_story/domain/category"
)

type Story struct {
	ID          uint64               `gorm:"column:id"`
	Title       string               `gorm:"column:title"`
	Description string               `gorm:"column:description"`
	Pages       []StoryPage          `gorm:"foreignKey:story_id;references:id"`
	ImageID     *uint64              `gorm:"column:image_id"`
	Image       *Image               `gorm:"foreignKey:image_id;references:id"`
	Categories  []dcategory.Category `gorm:"many2many:story_to_category;foreignKey:id;joinForeignKey:story_id;References:id;joinReferences:category_id;gorm:ordered"`
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}

type StoryPage struct {
	ID           uint64      `gorm:"column:id"`
	StoryID      uint64      `gorm:"column:story_id"`
	PageNumber   int         `gorm:"column:page_number"`
	ContentCN    string      `gorm:"column:content_cn"`
	ContentHakka string      `gorm:"column:content_hakka"`
	AudioFiles   []AudioFile `gorm:"foreignKey:story_page_id;references:id"`
}

type Image struct {
	ID       uint64 `gorm:"column:id"`
	ImageURL string `gorm:"column:image_url"`
}

type AudioFile struct {
	ID          uint64 `gorm:"column:id"`
	StoryPageID uint64 `gorm:"column:story_page_id"`
	Dialect     string `gorm:"column:dialect"`
	AudioURL    string `gorm:"column:audio_url"`
}

type StoryToCategory struct {
	StoryID    uint64 `gorm:"column:story_id"`
	CategoryID uint64 `gorm:"column:category_id"`
	CreatedAt  time.Time
}

func (Story) TableName() string {
	return "stories"
}

func (StoryPage) TableName() string {
	return "story_pages"
}

func (Image) TableName() string {
	return "images"
}

func (AudioFile) TableName() string {
	return "audio_files"
}

func (StoryToCategory) TableName() string {
	return "story_to_category"
}
