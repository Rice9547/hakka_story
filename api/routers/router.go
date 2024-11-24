package routers

import (
	"github.com/gin-gonic/gin"

	hauth "github.com/rice9547/hakka_story/api/handlers/auth"
	himage "github.com/rice9547/hakka_story/api/handlers/image"
	hstory "github.com/rice9547/hakka_story/api/handlers/story"
	"github.com/rice9547/hakka_story/lib/uploader"
	"github.com/rice9547/hakka_story/persistence/mysql"
	sstory "github.com/rice9547/hakka_story/service/story"
	supload "github.com/rice9547/hakka_story/service/upload"
)

func InitRoutes(apiRoute *gin.RouterGroup, adminRoute *gin.RouterGroup, db *mysql.Client, imageUploader *uploader.Client) {
	storyRepo := mysql.NewStory(db)
	storyService := sstory.New(storyRepo)
	imageService := supload.New(imageUploader)

	storyHandler := hstory.New(storyService)
	imageHandler := himage.New(imageService)

	apiRoute.GET("/story", storyHandler.List)
	apiRoute.GET("/story/:id", storyHandler.Get)

	adminRoute.GET("/auth", hauth.Auth)

	adminRoute.POST("/image/upload", imageHandler.Upload)

	adminStoryRoutes := adminRoute.Group("/story")
	adminStoryRoutes.POST("", storyHandler.Create)
	adminStoryRoutes.PUT("/:id", storyHandler.Update)
}
