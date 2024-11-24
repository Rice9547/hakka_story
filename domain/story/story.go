package dstory

import "gorm.io/gorm"

type Story struct {
	gorm.Model
	ID          uint64      `gorm:"column:id"`
	Title       string      `gorm:"column:title"`
	Description string      `gorm:"column:description"`
	Pages       []StoryPage `gorm:"foreignKey:story_id;references:id"`

	ImageID *uint64 `gorm:"column:image_id"`
	Image   *Image  `gorm:"foreignKey:image_id;references:id"`
}

type StoryPage struct {
	gorm.Model
	ID           uint64 `gorm:"column:id"`
	StoryID      uint64 `gorm:"column:story_id"`
	PageNumber   int    `gorm:"column:page_number"`
	ContentCN    string `gorm:"column:content_cn"`
	ContentHakka string `gorm:"column:content_hakka"`
}

type Image struct {
	ID       uint64 `gorm:"column:id"`
	ImageURL string `gorm:"column:image_url"`
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
