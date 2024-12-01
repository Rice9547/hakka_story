package supload

import (
	"context"
	"io"
	"mime/multipart"
	"path/filepath"

	"github.com/google/uuid"

	"github.com/rice9547/hakka_story/lib/errors"
	"github.com/rice9547/hakka_story/repository"
)

type UploadService struct {
	uploadRepo repository.Upload
}

func New(uploadRepo repository.Upload) *UploadService {
	return &UploadService{
		uploadRepo: uploadRepo,
	}
}

func (s *UploadService) UploadImage(ctx context.Context, file io.Reader, header *multipart.FileHeader) (string, error) {
	contentType := header.Header.Get("Content-Type")
	filename := uuid.New().String() + filepath.Ext(header.Filename)

	url, err := s.uploadRepo.UploadImage(ctx, file, filename, contentType)
	if err != nil {
		return "", errors.ErrFailedToUploadFile
	}

	return url, nil
}

func (s *UploadService) UploadAudio(ctx context.Context, file io.Reader, header *multipart.FileHeader) (string, error) {
	contentType := header.Header.Get("Content-Type")
	filename := uuid.New().String() + filepath.Ext(header.Filename)

	url, err := s.uploadRepo.UploadAudio(ctx, file, filename, contentType)
	if err != nil {
		return "", errors.ErrFailedToUploadFile
	}

	return url, nil
}
