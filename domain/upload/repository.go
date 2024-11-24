package dupload

import (
	"context"
	"io"
)

type Repository interface {
	UploadImage(ctx context.Context, file io.Reader, fileName, contentType string) (string, error)
	UploadAudio(ctx context.Context, file io.Reader, fileName, contentType string) (string, error)
}
