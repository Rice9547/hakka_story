package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/rice9547/hakka_story/api/middlewares"
	"github.com/rice9547/hakka_story/api/routers"
	_ "github.com/rice9547/hakka_story/docs"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/rice9547/hakka_story/config"
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

	mw := middlewares.NewAuthMiddlewares(cfg.Admin, cfg.Auth0)
	router := gin.Default()
	router.Use(middlewares.CORSMiddleware())
	apiRoute := router.Group("/api")
	adminRoute := apiRoute.Group("/admin")
	adminRoute.Use(mw.AuthMiddleware(), mw.AdminOnlyMiddleware())

	routers.InitRoutes(apiRoute, adminRoute, db)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Printf("Server is running on port %d\n", cfg.Server.Port)
	if err := router.Run(fmt.Sprintf(":%d", cfg.Server.Port)); err != nil {
		panic(fmt.Errorf("failed to start server: %v", err))
	}
}
