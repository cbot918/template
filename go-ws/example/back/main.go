package main

import (
	"fmt"
	"io"
	"net/http"

	"golang.org/x/net/websocket"
)

type Server struct {
	conns map[*websocket.Conn]bool
}

func NewServer() *Server {
	return &Server{
		conns: make(map[*websocket.Conn]bool),
	}
}

func (s *Server) handleWS(ws *websocket.Conn) {
	fmt.Println("new incoming connection from client: ", ws.RemoteAddr())

	s.conns[ws] = true

	s.readLoop(ws)
}

func (s *Server) readLoop(ws *websocket.Conn) {
	buf := make([]byte, 1024)
	for {
		n, err := ws.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("read error: ", err)
			continue
		}
		msg := buf[:n]
		s.broadcast(msg)
	}
}

func (s *Server) broadcast(b []byte) {
	for conn := range s.conns {
		go func(conn *websocket.Conn) {
			if _, err := conn.Write(b); err != nil {
				fmt.Println("write error: ", err)
			}
		}(conn)
	}
}

func main() {
	server := NewServer()
	http.Handle("/ws", websocket.Handler(server.handleWS))
	if err := http.ListenAndServe(":8886", nil); err != nil {
		panic(err)
	}
}
