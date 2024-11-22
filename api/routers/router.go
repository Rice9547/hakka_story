package routers

import (
	"github.com/gin-gonic/gin"

	hstory "github.com/rice9547/hakka_story/api/handlers/story"
	"github.com/rice9547/hakka_story/api/middlewares"
	"github.com/rice9547/hakka_story/persistence/mysql"
	sstory "github.com/rice9547/hakka_story/service/story"
)

func InitRoutes(router *gin.Engine, db *mysql.Client) {
	storyRepo := mysql.NewStory(db)
	storyService := sstory.New(storyRepo)

	storyHandler := hstory.New(storyService)

	api := router.Group("/api")
	api.GET("/story", storyHandler.List)
	api.GET("/story/:id", storyHandler.Get)

	adminRoutes := api.Group("/admin")
	adminRoutes.Use(middlewares.AuthMiddleware(), middlewares.AdminOnlyMiddleware())

	adminStoryRoutes := adminRoutes.Group("/story")
	adminStoryRoutes.POST("", storyHandler.Create)
	adminStoryRoutes.PUT("/:id", storyHandler.Update)
}
