package main

import (
	"gin/handlers"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	// Load HTML templates
	router.LoadHTMLGlob("templates/*")

	// Define the route using the handler from another file
	router.GET("/", handlers.Home)
	router.GET("/about",handlers.About)

	return router
}

func main() {
	router := setupRouter()

	// Start serving the application
	println("localhost:8080")
	router.Run(":8080") // listens on :8080 by default
}
