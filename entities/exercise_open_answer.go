package entities

type ExerciseOpenAnswer struct {
	ID         uint64 `gorm:"column:id"`
	ExerciseID uint64 `gorm:"column:exercise_id"`
	AnswerText string `gorm:"column:answer_text"`
}

func (ExerciseOpenAnswer) TableName() string {
	return "exercise_open_answers"
}
