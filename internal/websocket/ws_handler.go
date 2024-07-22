package websocket

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Handler struct {
	hub Hub
}

func NewHandler(h Hub) *Handler {
	return &Handler{
		hub: h,
	}
}

type CreateRoomReq struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (h Handler) CreateRoom(c *gin.Context) {
	var req CreateRoomReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	h.hub.Rooms[req.ID] = Room{
		ID:      req.ID,
		Name:    req.Name,
		Clients: make(map[string]*Client),
	}

	c.JSON(200, gin.H{"message": "room created"})
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *Handler) JoinRoom(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Failed to set websocket upgrade:", err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	roomID := c.Param("roomId")
	clientID := c.Query("userId")
	name := c.Query("name")

	cl := Client{
		Conn:    conn,
		ID:      clientID,
		Name:    name,
		RoomID:  roomID,
		Message: make(chan *Message, 10),
	}

	m := Message{
		Content: "Một người dùng đã tham gia đoạn chat",
		RoomID:  roomID,
		Name:    name,
	}

	fmt.Println("Joining room", roomID)
	// Register client
	h.hub.Register <- &cl
	log.Printf("Client registered: %s in room %s", clientID, roomID)
	// Broadcast message
	h.hub.Broadcast <- &m
	log.Printf("Message broadcasted: %s joined room %s", name, roomID)
	go cl.readMessage(&h.hub)
	go cl.writeMessage()
}

type RoomRes struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (h *Handler) GetRooms(c *gin.Context) {
	rooms := make([]RoomRes, 0)

	for _, r := range h.hub.Rooms {
		rooms = append(rooms, RoomRes{
			ID:   r.ID,
			Name: r.Name,
		})
	}

	c.JSON(200, gin.H{"rooms": rooms})
}

type ClientRes struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
func (h *Handler) GetClients(c *gin.Context){
	var clients []ClientRes

	roomID := c.Param("roomId")
	if _,ok := h.hub.Rooms[roomID]; !ok {
		clients = make([]ClientRes, 0)
		c.JSON(200, gin.H{"clients": clients})
	}

	for _, c := range h.hub.Rooms[roomID].Clients {
		clients = append(clients, ClientRes{
			ID:   c.ID,
			Name: c.Name,
		})
	}

	c.JSON(200, gin.H{"clients": clients})
}