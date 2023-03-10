package main

import (
	"example/rest-api/app/models"
	db "example/rest-api/config"
)

func main() {
	db := db.ConnectDB()
	db.AutoMigrate(&models.Album{})
	db.AutoMigrate(&models.Message{})
	db.AutoMigrate(&models.User{})
}
