package entities

type StoryPage struct {
	ID           uint64      `gorm:"column:id"`
	StoryID      uint64      `gorm:"column:story_id"`
	PageNumber   int         `gorm:"column:page_number"`
	ContentCN    string      `gorm:"column:content_cn"`
	ContentHakka string      `gorm:"column:content_hakka"`
	AudioFiles   []AudioFile `gorm:"foreignKey:story_page_id;references:id"`
	ImageID      *uint64     `gorm:"column:image_id"`
	Image        *Image      `gorm:"foreignKey:image_id;references:id;OnDelete:CASCADE"`
}

func (StoryPage) TableName() string {
	return "story_pages"
}
