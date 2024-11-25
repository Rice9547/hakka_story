package haudio

import supload "github.com/rice9547/hakka_story/service/upload"

type Audio struct {
	uploader  *supload.UploadService
	generator audioGenerator
}

type audioGenerator interface {
	Text2Speech(prompt string) ([]byte, error)
}

func New(uploader *supload.UploadService, generator audioGenerator) *Audio {
	return &Audio{
		uploader:  uploader,
		generator: generator,
	}
}
