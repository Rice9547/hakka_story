package entities

type Category struct {
	ID   uint64 `gorm:"column:id"`
	Name string `gorm:"column:name"`
}

func (Category) TableName() string {
	return "categories"
}
