package entities

type (
	ExerciseType int

	Exercise struct {
		ID         uint64       `gorm:"column:id"`
		StoryID    uint64       `gorm:"column:story_id"`
		Type       ExerciseType `gorm:"column:type"`
		PromptText string       `gorm:"column:prompt_text"`
		AudioURL   string       `gorm:"column:audio_url"`

		Choices []ExerciseChoice     `gorm:"foreignKey:exercise_id;references:id"`
		Answers []ExerciseOpenAnswer `gorm:"foreignKey:exercise_id;references:id"`
	}
)

const (
	ExerciseTypeFillIn ExerciseType = iota
	ExerciseTypeChoice
)

func (Exercise) TableName() string {
	return "exercises"
}
