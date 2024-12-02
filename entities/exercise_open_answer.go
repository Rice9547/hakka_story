package entities

import "time"

type ExerciseOpenAnswer struct {
	ID         int       `gorm:"column:id"`
	ExerciseID int       `gorm:"column:exercise_id"`
	Answer     string    `gorm:"column:answer"`
	CreatedAt  time.Time `gorm:"column:created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at"`
}

func (ExerciseOpenAnswer) TableName() string {
	return "exercise_open_answers"
}
