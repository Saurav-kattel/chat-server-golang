package main

import (
	"log"
	"net/http"

	"github.com/rs/cors"
	"x-clone.com/chat-server/internal/db"
	"x-clone.com/chat-server/internal/handlers"
	"x-clone.com/chat-server/internal/socket"
)

func main() {

	_ = socket.NewServer()
	db, err := db.ConnectDB("")

	if err != nil {
		log.Fatal(err)
	}
	routes := handlers.Route(db)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "auth_token_x_clone"},
		AllowCredentials: true,
		Debug:            false,
	})

	server := http.Server{
		Addr:    "localhost:5000",
		Handler: c.Handler(routes),
	}
	log.Fatal(server.ListenAndServe())
}
