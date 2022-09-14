package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

func listener(conn *websocket.Conn) {
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			logrus.Error("Error during message reading: ", err)
		}
		logrus.Infof("Received from %s: %s", conn.RemoteAddr().String(), message)
	}
}

func writer(conn *websocket.Conn) {
	var input string
	for {
		_, err := fmt.Scan(&input)
		if err != nil {
			logrus.Error("Error scan: ", err)
		}
		err = conn.WriteMessage(1, []byte(input))
		if err != nil {
			logrus.Error("Error during message writing: ", err)
		}
	}
}

func main() {
	socketUrl := "ws://localhost:7077" + "/ws"
	conn, _, err := websocket.DefaultDialer.Dial(socketUrl, nil)
	if err != nil {
		logrus.Fatal("Error connecting to Websocket Server:", err)
	}
	go listener(conn)
	writer(conn)
}
