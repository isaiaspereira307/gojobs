package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Initialize the default Gin router
	router := gin.Default()

	// Initialize the routes
	initializeRoutes(router)

	// Run the server
	router.Run(":8080")
}
