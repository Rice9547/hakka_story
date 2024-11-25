package himage

import (
	"github.com/rice9547/hakka_story/lib/openai"
	supload "github.com/rice9547/hakka_story/service/upload"
)

type Image struct {
	uploader  *supload.UploadService
	generator imageGenerator
}

type imageGenerator interface {
	Text2Image(prompt string) (string, error)
}

func New(uploader *supload.UploadService, generator *openai.Client) *Image {
	return &Image{
		uploader:  uploader,
		generator: generator,
	}
}
