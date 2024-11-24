package supload

import (
	"context"
	"mime/multipart"
	"path/filepath"

	"github.com/google/uuid"

	dupload "github.com/rice9547/hakka_story/domain/upload"
	"github.com/rice9547/hakka_story/lib/errors"
)

type UploadService struct {
	repo dupload.Repository
}

func New(repo dupload.Repository) *UploadService {
	return &UploadService{
		repo: repo,
	}
}

func (s *UploadService) UploadImage(ctx context.Context, file multipart.File, header *multipart.FileHeader) (string, error) {
	contentType := header.Header.Get("Content-Type")
	filename := uuid.New().String() + filepath.Ext(header.Filename)

	url, err := s.repo.UploadImage(ctx, file, filename, contentType)
	if err != nil {
		return "", errors.ErrFailedToUploadImage
	}

	return url, nil
}
