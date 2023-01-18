package controllers

import (
	"example/rest-api/app/models"
	"example/rest-api/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User models.User

func GetAllUsers(c *gin.Context) {
	var users []User
	db := config.ConnectDB()
	err := db.Find(&users)
	if err.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error.Error()})
	}
	c.IndentedJSON(http.StatusOK, users)
}
