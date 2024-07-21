package router

import (
	"golang-test/internal/user"
	"golang-test/internal/websocket"
	"golang-test/middleware"
	"log"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func InitRouter(userhandler user.Handler, wsHandler *websocket.Handler) {
	r = gin.Default()
	r.Use(middleware.CORSMiddleware())
	r.POST("/register", userhandler.CreateUser)
	r.POST("/login", userhandler.Login)
	r.POST("/logout", userhandler.Logout)

	// websocket
	r.POST("/ws/create-room", wsHandler.CreateRoom)
	r.GET("/ws/join-room/:roomId", wsHandler.JoinRoom)
	r.GET("/ws/get-rooms", wsHandler.GetRooms)
	r.GET("/ws/get-clients/:roomId", wsHandler.GetClients)
	log.Println("Router initialized")

}

func Run(addr string) error {
	log.Println("Server starting on", addr)
	return r.Run(addr)
}
