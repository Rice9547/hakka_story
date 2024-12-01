package middlewares

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware(allowOrigins []string) gin.HandlerFunc {
	c := cors.DefaultConfig()
	c.AllowCredentials = true
	c.AllowOrigins = allowOrigins
	c.AddAllowHeaders("Authorization")

	return cors.New(c)
}
