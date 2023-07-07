package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"
	"golang.org/x/net/websocket"
)

var log = fmt.Println

type Chaty struct {
	Conns  []Conn
	Counts int32
}

type Conn struct {
	Id   string
	Name string `json:"name"`
	Conn *websocket.Conn
	User User
}

type User struct {
	Name string `json:"_init_user_name"`
}

type Message struct {
	Content string `json:"content"`
	Group   string `json:"group"`
}

func NewChaty() *Chaty {
	return &Chaty{
		Counts: 0,
	}
}

func (c *Chaty) handleWs(ws *websocket.Conn) {
	c.Counts += 1
	conn := Conn{
		Id:   uuid.New().String(),
		Conn: ws,
	}
	c.Conns = append(c.Conns, conn)

	log("users: ", c.Counts)

	c.readListener(conn)

	// for _, conn := range c.Conns {

	// }

	// j, err := json.MarshalIndent(ws, "", "  ")
	// if err != nil {
	// 	log("marshal err")
	// 	return
	// }

	// log(string(j))
	// wsRead(ws)
}

func (c *Chaty) readListener(ws Conn) {
	buf := make([]byte, 1024)

	for {
		n, err := ws.Conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				log("client disconnect")
				break
			}
			log("ws read failed")
			continue
		}
		msg := buf[:n]

		if IsFirstMsg(msg) { // if init message with user name
			var user User
			err = json.Unmarshal(buf[:n], &user)
			if err != nil {
				log("marshal user failed")
				continue
			}
			ws.User.Name = user.Name

			log("id: ", ws.Id)
			log("name: ", ws.User.Name)
		} else { // else deal with other message with data and channel
			var message Message
			err = json.Unmarshal(msg, &message)
			if err != nil {
				log("marshal message failed")
				continue
			}
			log("group: ", message.Group)
			log("data: ", message.Content)
		}

	}

}

func main() {

	chat := NewChaty()

	http.Handle("/ws", websocket.Handler(chat.handleWs))
	http.ListenAndServe(":8887", nil)
}
