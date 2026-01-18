package server

import "github.com/gorilla/websocket"

type Client struct {
	Conn     *websocket.Conn
	Send     chan []byte
	Username string
}
