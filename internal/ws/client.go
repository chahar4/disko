package ws

import (
	"encoding/json"
	"log"

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
		c.Conn.Close()
		close(c.Send)
	}()
	for {
		_, msg, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}
		var input Message
		if err := json.Unmarshal(msg, &input); err != nil {
			log.Println(err)
			continue
		}
		if _, ok := c.Rooms[input.roomID]; !ok {
			log.Println("client is not in channel")
			continue
		}
		c.hub.Broadcast <- input
	}
}
