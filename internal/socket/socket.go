package socket

import (
	"io"
	"log"
	"sync"

	"golang.org/x/net/websocket"
)

type Server struct {
	conns map[*websocket.Conn]string
	rooms map[string][]*websocket.Conn
	mu    *sync.Mutex
}

// function that returns a Server
func NewServer() *Server {
	return &Server{
		conns: make(map[*websocket.Conn]string),
		rooms: make(map[string][]*websocket.Conn),
		mu:    &sync.Mutex{},
	}
}

func (s *Server) JoinRoom(ws *websocket.Conn, roomId string) {
	s.mu.Lock()
	s.conns[ws] = roomId
	s.rooms[roomId] = append(s.rooms[roomId], ws)
	s.mu.Unlock()
}

func (s *Server) LeaveRoom(ws *websocket.Conn) {
	s.mu.Lock()

	roomId, ok := s.conns[ws]
	if ok {
		conns := s.rooms[roomId]
		for i, conn := range conns {
			if conn == ws {
				s.rooms[roomId] = append(conns[:i], conns[i+1:]...)
				break
			}
		}
		delete(s.conns, ws)
	}
	s.mu.Unlock()
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
