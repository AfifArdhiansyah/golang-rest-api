package app

import (
	controllers "example/rest-api/app/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()

	getAlbum, err := controllers.GetAllAlbum()

	if err != nil {
		panic(err)
	}

	router.GET("/albums", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, getAlbum)
	})

	return router
}
