package controllers

import (
	models "example/rest-api/app/models"
	"example/rest-api/config"
	"fmt"
)

type Album models.Album

func GetAllAlbum() ([]Album, error) {
	var albums []Album

	db := config.ConnectDB()
	defer db.Close()

	// db.Query("SELECT * FROM album")
	getAllAlbum := "SELECT * FROM album"
	rows, err := db.Query(getAllAlbum)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("albums: %v", err)
		}
		albums = append(albums, alb)
	}

	return albums, nil
}
