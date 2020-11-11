package main

import (
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/handlers"
	"github.com/Max-Levitskiy/cloud-resource-dashboard/api/logger"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	runApp()
}

func runApp() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowHeaders:  []string{"Origin"},
		ExposeHeaders: []string{"Content-Length"},
	}))

	router.GET("/ping", handlers.StatusHandler)
	//router.GET("/resource", handlers.ResourceHandler)
	router.GET("/resource/count", handlers.ResourceCountHandler)
	router.GET("/resource/distinct/:field", handlers.ResourceDistinctServiceHandler)
	router.POST("/resource/scan/full", handlers.FullScanHandler)

	logger.Info.Print("Starting api server")
	if err := router.Run(":8080"); err != nil {
		logger.Error.Fatal(err)
	}
}
