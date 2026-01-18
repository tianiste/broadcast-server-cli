package server

import (
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

type Hub struct {
	Clients    map[*Client]bool
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan []byte
	Mu         sync.Mutex
}

func NewHub() *Hub {
	return &Hub{
		Clients:    make(map[*Client]bool),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan []byte),
		Mu:         sync.Mutex{},
	}
}

func (h *Hub) Run() {
	// something
}

func (h *Hub) HandleConn(w http.ResponseWriter, r *http.Request) {
	// something
}

func (h *Hub) StartServer(port string) {
	// something
}
