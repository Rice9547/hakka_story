package errors

import (
	"errors"
	"fmt"
)

var (
	ErrStoryNotFound    = errors.New("story not found")
	ErrCategoryNotFound = errors.New("category not found")
	ErrExerciseNotFound = errors.New("exercise not found")

	ErrUnsupportedFileType = errors.New("unsupported file type")
	ErrFailedToUploadFile  = errors.New("failed to upload file")

	ErrExerciseTypeMismatch   = errors.New("exercise type mismatch")
	ErrExerciseTypeNotSupport = errors.New("exercise type not supported")
	ErrExerciseInvalidAnswer  = errors.New("exercise invalid answer")

	ErrUnauthorized = errors.New("unauthorized")
)

func Is(err error, target error) bool {
	return errors.Is(err, target)
}

func NewUnauthorizedError(message string) error {
	return fmt.Errorf("%w: %s", ErrUnauthorized, message)
}
