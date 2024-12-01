package repository

import (
	"context"
	"io"
)

type Upload interface {
	UploadImage(ctx context.Context, file io.Reader, fileName, contentType string) (string, error)
	UploadAudio(ctx context.Context, file io.Reader, fileName, contentType string) (string, error)
}
