package websocket

import "fmt"

type Room struct {
	ID      string             `json:"id"`
	Name    string             `json:"name"`
	Clients map[string]*Client `json:"clients"`
}
type Hub struct {
	Rooms      map[string]Room
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan *Message
}

func NewHub() *Hub {
	return &Hub{
		Rooms:      make(map[string]Room),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan *Message),
	}
}

func (h Hub) Run() {
	fmt.Printf("Hub running\n")
	for {
		select {
		case cl := <-h.Register:
			room, exists := h.Rooms[cl.RoomID]
			if !exists {
				room = Room{
					ID:      cl.RoomID,
					Name:    cl.RoomID, // Or another way to set the name
					Clients: make(map[string]*Client),
				}
				h.Rooms[cl.RoomID] = room
			}
			room.Clients[cl.ID] = cl
		case cl := <-h.Unregister:
			if _, ok := h.Rooms[cl.RoomID]; ok {
				if _, ok := h.Rooms[cl.RoomID].Clients[cl.ID]; ok {
					if len(h.Rooms[cl.RoomID].Clients) != 0 {
						h.Broadcast <- &Message{
							Content: cl.Name + " left",
							RoomID:  cl.RoomID,
							Name:    cl.Name,
						}
					}

					delete(h.Rooms[cl.RoomID].Clients, cl.ID)
					close(cl.Message)
				}
			}
		case m := <-h.Broadcast:
			if _, ok := h.Rooms[m.RoomID]; ok {
				for _, cl := range h.Rooms[m.RoomID].Clients {
					cl.Message <- m
				}
			}
		}
	}
}
