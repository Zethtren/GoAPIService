package main

import (
	"example.com/handlers"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	// Create router
	router := gin.Default()

	// Connect functions to path
	router.GET("/albums", handlers.GetAlbums)
	router.GET("/albums/byArtist/:name", handlers.GetAlbumsByAuthorName)
	router.GET("/albums/lessThan/:price", handlers.GetAlbumsPriceLessThan)
	router.POST("/albums", handlers.InsertAlbum)

	// Start the server after passing all of its paths and arguments
	router.Run("localhost:8080")
}
