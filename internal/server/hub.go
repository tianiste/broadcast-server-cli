package server

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

const (
	ReadBufferSize  = 1024
	WriteBufferSize = 1024
)

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  ReadBufferSize,
	WriteBufferSize: WriteBufferSize,
}

type Server struct {
	Clients map[*Client]bool

	Join chan *Client

	Leave chan *Client

	Forward chan []byte

	Mu sync.Mutex
}

func NewServer() *Server {
	return &Server{
		Clients: make(map[*Client]bool),
		Join:    make(chan *Client),
		Leave:   make(chan *Client),
		Forward: make(chan []byte),
		Mu:      sync.Mutex{},
	}
}

func (s *Server) Run() {
	for {
		select {
		case client := <-s.Join:

			s.Mu.Lock()
			s.Clients[client] = true
			s.Mu.Unlock()

		case client := <-s.Leave:
			s.Mu.Lock()
			delete(s.Clients, client)
			s.Mu.Unlock()
			s.Forward <- []byte(fmt.Sprintf("%s left the chat", client.Username))

		case msg := <-s.Forward:
			s.Mu.Lock()
			clients := make([]*Client, 0, len(s.Clients))
			for c := range s.Clients {
				clients = append(clients, c)
			}
			s.Mu.Unlock()
			for client := range s.Clients {
				if err := client.Socket.WriteMessage(websocket.TextMessage, msg); err != nil {
					fmt.Printf("error writing message: %v\n", err)
					return
				}
			}
		}
	}
}

func (s *Server) HandleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("Error upgrading connection: %v\n", err)
		return
	}

	client := &Client{
		Socket:   conn,
		Recieve:  make(chan []byte),
		Username: r.PathValue("username"),
	}

	s.Forward <- []byte(fmt.Sprintf("%s joined the chat", client.Username))
	s.Join <- client

	defer func() {
		client.Socket.Close()
		s.Leave <- client
	}()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Error reading message: %v\n", err)
			break
		}

		s.Forward <- []byte(fmt.Sprintf("%s:%s\n", client.Username, string(msg)))
	}

}

func StartServer(port string) {
	addr := fmt.Sprintf(":%s", port)

	server := NewServer()

	http.HandleFunc("/ws/{username}", server.HandleConnections)

	go server.Run()

	log.Printf("Starting websocket server on port %s", port)
	log.Fatal(http.ListenAndServe(addr, nil))
}
