package entities

import "time"

type (
	ExerciseType int

	Exercise struct {
		ID         int          `gorm:"column:id"`
		StoryID    int          `gorm:"column:story_id"`
		Type       ExerciseType `gorm:"column:type"`
		PromptText string       `gorm:"column:prompt_text"`
		AudioURL   string       `gorm:"column:audio_url"`
		CreatedAt  time.Time    `gorm:"column:created_at"`
		UpdatedAt  time.Time    `gorm:"column:updated_at"`

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
