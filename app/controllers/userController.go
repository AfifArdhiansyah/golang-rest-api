package controllers

import (
	"example/rest-api/app/middlewares"
	"example/rest-api/app/models"
	"example/rest-api/config"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type User models.User
type RequestUser models.RequestUser

// encrypt password
func (u *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	u.Password = string(bytes)
	return nil

}

// check password
func (u *User) CheckPasswordHash(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return err
	}
	return nil
}

func GetAllUsers(c *gin.Context) {
	var users []User
	db := config.ConnectDB()
	err := db.Find(&users)
	if err.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error.Error()})
	}
	c.IndentedJSON(http.StatusOK, users)
}

func Register(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := user.HashPassword(user.Password); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	db := config.ConnectDB()
	err := db.Create(&user)
	if err.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error.Error()})
	}
	c.IndentedJSON(http.StatusOK, user)
}

func Login(c *gin.Context) {
	var request RequestUser
	var user User
	if err := c.ShouldBindJSON(&request); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	db := config.ConnectDB()
	record := db.Where("username = ?", request.Username).First(&user)
	if record.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": record.Error.Error()})
	}
	if err := user.CheckPasswordHash(request.Password); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	tokenString, err := middlewares.GenerateToken(user.Username, user.Password)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{
		"data":  user,
		"token": tokenString,
	})
}
