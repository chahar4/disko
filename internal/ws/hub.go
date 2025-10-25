package ws

type Hub struct {
	Clients    map[string]*Client
	Rooms      map[string]map[*Client]struct{}
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan Message
}

func (h *Hub) Run() {
	for {
		select {
		case cli := <-h.Register:
			h.Clients[cli.ID] = cli
			for roomID := range cli.Rooms {
				if h.Rooms[roomID] == nil {
					h.Rooms[roomID] = make(map[*Client]struct{})
				}
				h.Rooms[roomID][cli] = struct{}{}

			}

		case cli := <-h.Unregister:
			delete(h.Clients, cli.ID)
			for roomID := range h.Rooms {
				delete(h.Rooms[roomID], cli)
				if len(h.Rooms[roomID]) == 0 {
					delete(h.Rooms, roomID)
				}

			}
			close(cli.Send)

		case msg := <-h.Broadcast:
			if _, ok := h.Rooms[msg.roomID]; ok {
				for cli := range h.Rooms[msg.roomID] {
					cli.Send <- msg.Payload
				}

			}

		}
	}
}
