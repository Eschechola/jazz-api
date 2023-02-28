package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"jazz-api.com/src/jazz-api.com/entities"
	repositories "jazz-api.com/src/jazz-api.com/repositories"
)

func main() {
	var router = gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", addAlbum)

	router.Run("localhost:8080")
}

func getAlbums(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, repositories.GetAlbums())
}

func addAlbum(context *gin.Context) {
	var newAlbum entities.Album

	if err := context.BindJSON(newAlbum); err != nil {
		return
	}

	var albumCreated = repositories.AddAlbum(newAlbum)
	context.IndentedJSON(http.StatusCreated, albumCreated)
}
