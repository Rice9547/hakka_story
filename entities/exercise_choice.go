package entities

type ExerciseChoice struct {
	ID         uint64 `gorm:"column:id"`
	ExerciseID uint64 `gorm:"column:exercise_id"`
	ChoiceText string `gorm:"column:choice_text"`
	IsCorrect  bool   `gorm:"column:is_correct"`
}

func (ExerciseChoice) TableName() string {
	return "exercise_choices"
}
