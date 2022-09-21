package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

func listener(conn *websocket.Conn) {
	for {
		_, username, err := conn.ReadMessage()
		_, message, err := conn.ReadMessage()
		if err != nil {
			logrus.Error("Error during message reading: ", err)
		}
		logrus.Infof("Received from %s: %s", username, message)
	}
}

func writer(conn *websocket.Conn) {
	var username string
	var input string
	for {
		_, err := fmt.Scanf("%s %s", &username, &input)
		if err != nil {
			logrus.Error("Error scan: ", err)
		}
		err = conn.WriteMessage(1, []byte(username))
		err = conn.WriteMessage(1, []byte(input))
		if err != nil {
			logrus.Error("Error during message writing: ", err)
		}
	}
}

func main() {
	var username string
	for {
		fmt.Print("Please enter your username: ")
		_, err := fmt.Scan(&username)
		if err != nil {
			logrus.Error("Wrong username: ", err)
		} else {
			break
		}
	}

	socketUrl := "ws://localhost:7077" + "/ws"
	conn, _, err := websocket.DefaultDialer.Dial(socketUrl, nil)
	err = conn.WriteMessage(1, []byte(username))
	if err != nil {
		logrus.Error("Error during username writing: ", err)
	}
	if err != nil {
		logrus.Fatal("Error connecting to Websocket Server:", err)
	}
	go listener(conn)
	writer(conn)
}
