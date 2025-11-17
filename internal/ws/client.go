package ws

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	ID    string
	Conn  *websocket.Conn
	Rooms map[string]struct{}
	Send  chan []byte
	hub   *Hub
}

type Message struct {
	roomID  string
	Payload []byte
}

func (c *Client) ReadMessage() {
	defer func() {
		c.Conn.Close()
	}()

	for {
		msg, ok := <-c.Send
		if !ok {
			return
		}

		err := c.Conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			return
		}
	}
}
