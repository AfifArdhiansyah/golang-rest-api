package app

import (
	controllers "example/rest-api/app/controllers"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()

	router.GET("/albums", controllers.GetAllAlbum)
	router.GET("/albums/:id", controllers.GetAlbum)
	router.POST("/albums", controllers.CreateAlbum)
	router.PUT("/albums/:id", controllers.UpdateAlbum)
	router.DELETE("/albums/:id", controllers.DeleteAlbum)

	router.GET("/ws", func(c *gin.Context) {
		controllers.WsEndpoint(c.Writer, c.Request)
	})
	router.GET("/messages", controllers.GetAllMessage)

	router.GET("/users", controllers.GetAllUsers)
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)

	return router
}
