package main

import (
	app "example/rest-api/app"
)

func main() {
	router := app.Router()

	router.Run("localhost:8080")
}
