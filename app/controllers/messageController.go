package controllers

import (
	models "example/rest-api/app/models"
	"example/rest-api/config"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Message models.Message

var (
	wsUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	wsConn *websocket.Conn
)

func WsEndpoint(w http.ResponseWriter, r *http.Request) {
	wsUpgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("Failed to set websocket upgrade: %+v", err)
		return
	}
	wsConn = ws
	defer wsConn.Close()

	for {
		var message Message
		err := wsConn.ReadJSON(&message)
		if err != nil {
			fmt.Printf("Failed to read JSON: %+v", err)
			continue
		}

		db := config.ConnectDB()
		result := db.Create(&message)

		if result.Error != nil {
			fmt.Printf("Failed to save message: %+v", result.Error)
			continue
		}

		fmt.Printf("Message received: %+v", message.Message)
	}
}

func GetAllMessage(c *gin.Context) {
	var messages []Message

	db := config.ConnectDB()

	result := db.Find(&messages)

	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": result.Error.Error()})
	}

	c.IndentedJSON(http.StatusOK, messages)
}
