package entities

type AudioFile struct {
	ID          uint64 `gorm:"column:id"`
	StoryPageID uint64 `gorm:"column:story_page_id"`
	Dialect     string `gorm:"column:dialect"`
	AudioURL    string `gorm:"column:audio_url"`
}

func (AudioFile) TableName() string {
	return "audio_files"
}
