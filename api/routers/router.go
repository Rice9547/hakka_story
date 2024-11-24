package routers

import (
	"github.com/gin-gonic/gin"

	hauth "github.com/rice9547/hakka_story/api/handlers/auth"
	hstory "github.com/rice9547/hakka_story/api/handlers/story"
	"github.com/rice9547/hakka_story/persistence/mysql"
	sstory "github.com/rice9547/hakka_story/service/story"
)

func InitRoutes(apiRoute *gin.RouterGroup, adminRoute *gin.RouterGroup, db *mysql.Client) {
	storyRepo := mysql.NewStory(db)
	storyService := sstory.New(storyRepo)

	storyHandler := hstory.New(storyService)

	apiRoute.GET("/story", storyHandler.List)
	apiRoute.GET("/story/:id", storyHandler.Get)

	adminRoute.GET("/auth", hauth.Auth)

	adminStoryRoutes := adminRoute.Group("/story")
	adminStoryRoutes.POST("", storyHandler.Create)
	adminStoryRoutes.PUT("/:id", storyHandler.Update)
}
