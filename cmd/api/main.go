package main

import (
	"fmt"
	"golang-test/db"
	"golang-test/internal/user"
	"golang-test/internal/websocket"
	"golang-test/router"
	"log"
)

func main() {
	fmt.Println("Starting application")
	client, err := db.ConnectMongo()
	if err != nil {
		log.Fatal(err)
	}

	userRepo := user.NewRepository(*db.GetCollection(client, "users").Database())
	userSrv := user.NewService(userRepo)
	userHandler := user.NewHandler(userSrv)
	hub := websocket.NewHub()
	wsHandler := websocket.NewHandler(*hub)

	go hub.Run()
	router.InitRouter(userHandler, wsHandler)

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to run server: ", err)
	}
}
