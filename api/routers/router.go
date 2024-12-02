package routers

import (
	"github.com/gin-gonic/gin"
	hexercise "github.com/rice9547/hakka_story/api/handlers/exercise"
	sexercise "github.com/rice9547/hakka_story/service/exercise"

	haudio "github.com/rice9547/hakka_story/api/handlers/audio"
	hauth "github.com/rice9547/hakka_story/api/handlers/auth"
	hcategory "github.com/rice9547/hakka_story/api/handlers/category"
	himage "github.com/rice9547/hakka_story/api/handlers/image"
	hstory "github.com/rice9547/hakka_story/api/handlers/story"
	htranslate "github.com/rice9547/hakka_story/api/handlers/translate"
	"github.com/rice9547/hakka_story/lib/openai"
	"github.com/rice9547/hakka_story/lib/uploader"
	"github.com/rice9547/hakka_story/persistence/mysql"
	scategory "github.com/rice9547/hakka_story/service/category"
	sstory "github.com/rice9547/hakka_story/service/story"
	supload "github.com/rice9547/hakka_story/service/upload"
)

func InitRoutes(
	apiRoute *gin.RouterGroup,
	adminRoute *gin.RouterGroup,
	db *mysql.Client,
	uploader *uploader.Client,
	openaiClient *openai.Client,
) {
	storyRepo := mysql.NewStory(db)
	categoryRepo := mysql.NewCategory(db)
	exerciseRepo := mysql.NewExercise(db)

	storyService := sstory.New(storyRepo)
	uploadService := supload.New(uploader)
	categoryService := scategory.New(categoryRepo)
	exerciseService := sexercise.New(exerciseRepo)

	storyHandler := hstory.New(storyService)
	categoryHandler := hcategory.New(categoryService)
	exerciseHandler := hexercise.New(exerciseService)
	imageHandler := himage.New(uploadService, openaiClient)
	audioHandler := haudio.New(uploadService, openaiClient)
	translateHandler := htranslate.New(openaiClient)

	apiRoute.GET("/story", storyHandler.List)
	apiRoute.GET("/story/:id", storyHandler.Get)
	apiRoute.GET("/category", categoryHandler.List)

	adminRoute.GET("/auth", hauth.Auth)

	adminRoute.POST("/image/upload", imageHandler.Upload)
	adminRoute.POST("/image/generate", imageHandler.Generate)
	adminRoute.POST("/audio/upload", audioHandler.Upload)
	adminRoute.POST("/audio/generate", audioHandler.Generate)
	adminRoute.POST("/translate/hakka", translateHandler.TranslateHakka)

	adminStoryRoutes := adminRoute.Group("/story")
	adminStoryRoutes.POST("", storyHandler.Create)
	adminStoryRoutes.PUT("/:id", storyHandler.Update)
	adminStoryRoutes.DELETE("/:id", storyHandler.Delete)
	adminStoryRoutes.GET("/exercise", exerciseHandler.CountStoriesExercise)

	adminExerciseRoutes := adminStoryRoutes.Group("/:id/exercise")
	adminExerciseRoutes.GET("", exerciseHandler.ListExerciseByStoryID)
	adminExerciseRoutes.POST("", exerciseHandler.CreateExercise)
	adminExerciseRoutes.PUT("/:exercise_id", exerciseHandler.UpdateExercise)

	adminCategoryRoutes := adminRoute.Group("/category")
	adminCategoryRoutes.POST("", categoryHandler.Create)
	adminCategoryRoutes.PUT("/:id", categoryHandler.Update)
	adminCategoryRoutes.DELETE("/:id", categoryHandler.Delete)
}
