package main

import (
	"log"
	"net/http"

	"golang.org/x/net/websocket"
	"x-clone.com/ws-server/db"
	"x-clone.com/ws-server/socket"
)

func main() {

	server := socket.NewServer()
	_, err := db.ConnectDB("")
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/ws", websocket.Handler(server.HandleWs))

	log.Fatal(http.ListenAndServe(":5000", nil))
}
