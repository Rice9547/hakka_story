package entities

type Image struct {
	ID       uint64 `gorm:"column:id"`
	ImageURL string `gorm:"column:image_url"`
}

func (Image) TableName() string {
	return "images"
}
