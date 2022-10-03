package main

import (
	"bufio"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"os"
)

func listener(conn *websocket.Conn) {
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			logrus.Error("Error during message reading: ", err)
		}
		logrus.Infof("Received from %s\n", message)
	}
}

func ponger(conn *websocket.Conn) {
	for {
		if err := conn.WriteMessage(websocket.PingMessage, nil); err != nil {
			return
		}
	}
}

func writer(conn *websocket.Conn) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		input := scanner.Text()
		//logrus.Info(input)
		if err := scanner.Err(); err != nil {
			logrus.Error("Error scan: ", err)
			break
		}
		err := conn.WriteMessage(1, []byte(input))
		if err != nil {
			logrus.Error("Error during message writing: ", err)
		}

	}
}

func main() {
	var username string
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Please enter your username: ")
		scanner.Scan()
		username = scanner.Text()
		if err := scanner.Err(); err != nil {
			logrus.Error("Error scan: ", err)
		} else {
			break
		}
	}

	socketUrl := "ws://localhost:7078" + "/ws"
	conn, _, err := websocket.DefaultDialer.Dial(socketUrl, nil)
	if err != nil {
		logrus.Fatal("Error connecting to Websocket Server:", err)
	}
	err = conn.WriteMessage(1, []byte(username))
	if err != nil {
		logrus.Error("Error during username writing: ", err)
	}

	//go ponger(conn)
	go listener(conn)
	writer(conn)
}
