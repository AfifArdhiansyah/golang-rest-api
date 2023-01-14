package controllers

import (
	models "example/rest-api/app/models"
	"example/rest-api/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Album models.Album

func GetAllAlbum(c *gin.Context) {
	var albums []Album

	db := config.ConnectDB()

	result := db.Find(&albums)

	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": result.Error.Error()})
	}

	c.IndentedJSON(http.StatusOK, albums)
}

func GetAlbum(c *gin.Context) {
	db := config.ConnectDB()
	id := c.Param("id")

	var album Album

	if err := db.Model(&album).Where("id = ?", id).First(&album).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	c.IndentedJSON(http.StatusOK, album)
}

func CreateAlbum(c *gin.Context) {
	db := config.ConnectDB()

	var newAlbum Album

	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	result := db.Create(&newAlbum)

	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": result.Error.Error()})
	}

	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func UpdateAlbum(c *gin.Context) {
	db := config.ConnectDB()
	id := c.Param("id")

	var album Album

	if err := db.Model(&album).Where("id = ?", id).First(&album).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	result := db.Model(&album).Where("id = ?", id).Updates(album)

	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": result.Error.Error()})
	}

	c.IndentedJSON(http.StatusOK, album)
}

func DeleteAlbum(c *gin.Context) {
	db := config.ConnectDB()
	id := c.Param("id")

	var album Album

	if err := db.Model(&album).Where("id = ?", id).First(&album).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	result := db.Delete(&album)

	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": result.Error.Error()})
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Album deleted successfully"})
}
