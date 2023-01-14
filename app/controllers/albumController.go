package controllers

import (
	models "example/rest-api/app/models"
	"example/rest-api/config"
)

type Album models.Album

func GetAllAlbum() ([]Album, error) {
	var albums []Album

	db := config.ConnectDB()

	result := db.Find(&albums)

	if result.Error != nil {
		return nil, result.Error
	}

	return albums, nil
}
