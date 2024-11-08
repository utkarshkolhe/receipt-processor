package main

import (
	"log"
	"net/http"
	"utkarsh/Fetch/api/routes"
	"utkarsh/Fetch/api/logger"
	"utkarsh/Fetch/api/config"
	"github.com/gin-gonic/gin"
)


// Entry point of the whole program
func main() {

  // Use the logger
  logger.Instance.Info(config.InfoApplicationStart)

	// Create a gin router to handle requests
	router := gin.Default()
	// Define all API routes
	routes.SetupRoutes(router)
	logger.Instance.Info(config.InfoRoutesStart)

	// Start the server on 8080. Port number can be changed in config.
	if err := http.ListenAndServe(":"+config.ServerPort, router); err != nil {
		log.Fatalf(config.ErrServerStartFailed+": %v", err)
	}
}
