package himage

import (
	"context"
	"io"
	"mime/multipart"
)

type (
	UploadService interface {
		UploadImage(ctx context.Context, file io.Reader, header *multipart.FileHeader) (string, error)
		UploadAudio(ctx context.Context, file io.Reader, header *multipart.FileHeader) (string, error)
	}

	ImageGenerator interface {
		Text2Image(ctx context.Context, prompt string) (string, []byte, error)
	}

	Image struct {
		uploader  UploadService
		generator ImageGenerator
	}
)

func New(uploader UploadService, generator ImageGenerator) *Image {
	return &Image{
		uploader:  uploader,
		generator: generator,
	}
}
