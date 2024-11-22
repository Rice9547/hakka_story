package response

import (
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

func Error(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, ResponseBase{
		Status:  "error",
		Message: message,
	})
}

func BadRequest(c *gin.Context, message string) {
	Error(c, http.StatusBadRequest, message)
}
