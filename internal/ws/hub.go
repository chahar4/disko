package ws

import "github.com/gorilla/websocket"

type Client struct {
	ID   int
	Conn *websocket.Conn
	Send chan []byte
}

type Hub struct {
}
