package response

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseBase struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type Response struct {
	ResponseBase
	Data interface{} `json:"data,omitempty"`
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		ResponseBase: ResponseBase{
			Status:  "success",
			Message: "OK",
		},
		Data: data,
	})
}

func Error(c *gin.Context, statusCode int, err error, message string) {
	log.Printf("Status: %d, Error: %v, Message: %s", statusCode, err, message)

	c.JSON(statusCode, ResponseBase{
		Status:  "error",
		Message: message,
	})
}

func BadRequest(c *gin.Context, err error, message string) {
	Error(c, http.StatusBadRequest, err, message)
}

func NotFound(c *gin.Context, message string) {
	Error(c, http.StatusNotFound, nil, message)
}

func InternalServerError(c *gin.Context, err error, message string) {
	Error(c, http.StatusInternalServerError, err, message)
}

func Unauthorized(c *gin.Context, message string) {
	Error(c, http.StatusUnauthorized, nil, message)
}

func Forbidden(c *gin.Context, message string) {
	Error(c, http.StatusForbidden, nil, message)
}
