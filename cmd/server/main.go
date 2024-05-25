package main

import (
	"log"
	"net/http"

	"x-clone.com/chat-server/internal/db"
	"x-clone.com/chat-server/internal/socket"
)

func main() {

	_ = socket.NewServer()
	_, err := db.ConnectDB("")
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(http.ListenAndServe(":5000", nil))
}
