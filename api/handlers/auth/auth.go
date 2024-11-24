package hauth

import (
	"github.com/gin-gonic/gin"

	"github.com/rice9547/hakka_story/lib/response"
)

type AuthResponse struct {
	IsAdmin bool `json:"isAdmin"`
}

// Auth godoc
// @Summary Check if the user is an admin
// @Description Get the admin status of the user
// @Tags admin auth
// @Accept  json
// @Produce  json
// @Success 200 {object} AuthResponse
// @Router /admin/auth [get]
func Auth(c *gin.Context) {
	response.Success(c, AuthResponse{IsAdmin: true})
}
