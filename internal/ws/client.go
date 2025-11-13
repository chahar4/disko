package ws

import (
	"encoding/json"

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

func (c *Client) WriteMessage() {
	defer func() {
		c.hub.Unregister <- c
		c.Conn.Close()
		close(c.Send)
	}()
	for {
		_, msg, err := c.Conn.ReadMessage()
		if err != nil {
			break
		}
		var input Message
		if err := json.Unmarshal(msg, &input); err != nil {
			continue
		}
		if _, ok := c.Rooms[input.roomID]; !ok {
			continue
		}
		c.hub.Broadcast <- &input
	}
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

		c.Conn.WriteJSON(msg)
	}
}
