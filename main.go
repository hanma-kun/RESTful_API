package main

import (
	"github.com/hanma-kun/REST/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Register the handlers
	router.GET("/albums", handlers.GetAlbums)
	router.GET("/albums/:id", handlers.GetAlbumByID)
	router.POST("/albums", handlers.PostAlbums)

	router.Run("localhost:8080")
}
