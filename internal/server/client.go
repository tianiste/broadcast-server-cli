package server

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/gorilla/websocket"
)

type Client struct {
	Socket   *websocket.Conn
	Recieve  chan []byte
	Username string
}

func Read(conn *websocket.Conn) {
	defer conn.Close()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("unable to read message", err)
			os.Exit(0)
			return
		}
		fmt.Println(string(msg))
	}
}

func Write(conn *websocket.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		if scanner.Scan() {
			msg := scanner.Text()
			if err := conn.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
				log.Println("unable to send message", err)
				return
			}
		} else if err := scanner.Err(); err != nil {
			log.Println("error reading input", err)
			os.Exit(0)
			return
		}
	}
}

func StartClient(port string, username string) {
	serverURL := fmt.Sprintf("ws://localhost:%s/ws/%s", port, username)
	conn, _, err := websocket.DefaultDialer.Dial(serverURL, nil)
	if err != nil {
		fmt.Println("error while connecting to server", err)
		return
	}
	defer conn.Close()
	go Read(conn)
	Write(conn)
}
