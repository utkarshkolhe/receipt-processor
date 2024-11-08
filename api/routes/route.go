package routes

import (
	"utkarsh/Fetch/api/handlers"

	"github.com/gin-gonic/gin"
)

// SetupRoutes sets up all API routes
func SetupRoutes(router *gin.Engine) {
	// API for adding reciept and getting associated ID
	router.POST("/receipts/process", handlers.ProcessReciept)

	// API to get points of a reciept using ID
	router.GET("/receipts/:id/points", handlers.GetPoints)

}
