package socket

import (
	"io"
	"log"
	"sync"

	"golang.org/x/net/websocket"
)

type Server struct {
	conns map[*websocket.Conn]bool
	mu    *sync.Mutex
}

// function that returns a Server
func NewServer() *Server {
	return &Server{
		conns: make(map[*websocket.Conn]bool),
		mu:    &sync.Mutex{},
	}
}

func (s *Server) ReadLoop(ws *websocket.Conn) {
	buf := make([]byte, 1024)
	for {
		n, err := ws.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Println("read error: ", err)
			continue
		}

		msg := buf[:n]
		log.Println(string(msg))
		ws.Write([]byte("Thank you dost"))
	}
}
