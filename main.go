package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/rice9547/hakka_story/api/middlewares"
	"github.com/rice9547/hakka_story/api/routers"
	"github.com/rice9547/hakka_story/config"
	_ "github.com/rice9547/hakka_story/docs"
	"github.com/rice9547/hakka_story/lib/openai"
	"github.com/rice9547/hakka_story/lib/uploader"
	"github.com/rice9547/hakka_story/persistence/mysql"
)

func main() {
	configPath := flag.String("config", "./config/config.yaml", "Path to the config file")
	flag.Parse()

	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		panic(fmt.Errorf("failed to load configuration: %v", err))
	}

	db, err := mysql.New(cfg.Database)
	if err != nil {
		panic(fmt.Errorf("failed to connect database: %v", err))
	}
	defer db.Close()

	s3Client, err := uploader.NewS3Client(cfg.Upload)
	if err != nil {
		panic(fmt.Errorf("failed to create S3 client: %v", err))
	}
	uploader := uploader.New(cfg.Upload, s3Client)

	openaiClient := openai.New(cfg.OpenAI)

	mw := middlewares.NewAuthMiddlewares(cfg.Auth0)
	router := gin.Default()
	router.Use(middlewares.CORSMiddleware(cfg.Server.AllowOrigins))
	apiRoute := router.Group("/api")
	adminRoute := apiRoute.Group("/admin")
	adminRoute.Use(mw.AuthMiddleware(true), mw.AdminOnlyMiddleware())

	routers.InitRoutes(apiRoute, adminRoute, db, uploader, openaiClient)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Printf("Server is running on port %d\n", cfg.Server.Port)
	if err := router.Run(fmt.Sprintf(":%d", cfg.Server.Port)); err != nil {
		panic(fmt.Errorf("failed to start server: %v", err))
	}
}
