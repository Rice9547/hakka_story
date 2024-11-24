package errors

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/rice9547/hakka_story/lib/response"
)

var (
	ErrInvalidInput = errors.New("invalid input")
	ErrInternal     = errors.New("internal server error")

	ErrStoryNotFound = errors.New("story not found")

	ErrUnsupportedFileType = errors.New("unsupported file type")
	ErrFailedToUploadFile  = errors.New("failed to upload file")
)

type AppError struct {
	StatusCode int
	Err        error
	Message    string
}

func (e *AppError) Error() string {
	return e.Err.Error()
}

func ErrorHandler(c *gin.Context, err error) {
	var appErr *AppError

	if errors.As(err, &appErr) {
		response.Error(c, appErr.StatusCode, appErr.Message)
	} else {
		response.Error(c, http.StatusInternalServerError, ErrInternal.Error())
	}
}

func NewAppError(statusCode int, err error, message string) *AppError {
	return &AppError{
		StatusCode: statusCode,
		Err:        err,
		Message:    message,
	}
}

func Is(err error, target error) bool {
	var appErr *AppError

	if errors.As(err, &appErr) {
		return errors.Is(appErr.Err, target)
	}

	return errors.Is(err, target)
}
