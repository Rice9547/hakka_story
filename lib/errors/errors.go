package errors

import (
	"errors"
)

var (
	ErrStoryNotFound    = errors.New("story not found")
	ErrCategoryNotFound = errors.New("category not found")
	ErrExerciseNotFound = errors.New("exercise not found")

	ErrUnsupportedFileType = errors.New("unsupported file type")
	ErrFailedToUploadFile  = errors.New("failed to upload file")
)

func Is(err error, target error) bool {
	return errors.Is(err, target)
}
