package entities

import (
	"gorm.io/gorm"
	"time"
)

type Story struct {
	ID          uint64      `gorm:"column:id"`
	Title       string      `gorm:"column:title"`
	Description string      `gorm:"column:description"`
	Pages       []StoryPage `gorm:"foreignKey:story_id;references:id"`
	ImageID     *uint64     `gorm:"column:image_id"`
	Image       *Image      `gorm:"foreignKey:image_id;references:id;OnDelete:CASCADE"`
	Categories  []Category  `gorm:"many2many:story_to_category;foreignKey:id;joinForeignKey:story_id;References:id;joinReferences:category_id;gorm:ordered"`
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}

func (Story) TableName() string {
	return "stories"
}
