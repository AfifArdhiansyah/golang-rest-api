package app

import (
	controllers "example/rest-api/app/controllers"
	middlewares "example/rest-api/app/middlewares"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "Hello World!"})
	})

	router.GET("/albums", controllers.GetAllAlbum)
	router.GET("/albums/:id", controllers.GetAlbum)
	router.POST("/albums", controllers.CreateAlbum)
	router.PUT("/albums/:id", controllers.UpdateAlbum)
	router.DELETE("/albums/:id", controllers.DeleteAlbum)

	router.GET("/ws", func(c *gin.Context) {
		controllers.WsEndpoint(c.Writer, c.Request)
	})

	router.GET("/users", controllers.GetAllUsers)
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)

	router.GET("/messages", controllers.GetAllMessage)
	router.Use(middlewares.Auth())
	{
		router.GET("/message", controllers.GetMessageByUser)
		router.POST("/message", controllers.SendMessage)

		router.GET("/profile", controllers.GetProfile)
	}

	return router
}
