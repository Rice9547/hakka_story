package entities

import "time"

type ExerciseOpenAnswer struct {
	ID         uint64    `gorm:"column:id"`
	ExerciseID uint64    `gorm:"column:exercise_id"`
	AnswerText string    `gorm:"column:answer_text"`
	CreatedAt  time.Time `gorm:"column:created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at"`
}

func (ExerciseOpenAnswer) TableName() string {
	return "exercise_open_answers"
}
