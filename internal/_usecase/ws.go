package _usecase

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type Client struct {
	Conn     *websocket.Conn
	Message  chan *Message
	ID       string `json:"id"`
	RoomID   string `json:"room_id"`
	Username string `json:"username"`
}

type Message struct {
	Content  string `json:"content"`
	RoomID   string `json:"room_id"`
	Username string `json:"username"`
}

type Room struct {
	ID      string             `json:"id"`
	Name    string             `json:"name"`
	Clients map[string]*Client `json:"clients"`
}

type RequestCreateRoom struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ClientResponse struct {
	ID       string `json:"id"`
	RoomID   string `json:"room_id"`
	Username string `json:"username"`
}

type RoomResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// origin := r.Header.Get("Origin")
		//
		// return origin == "http://localhost:3000"
		return true
	},
}

func (c *Client) WriteMessage() {
	defer func() {
		c.Conn.Close()
	}()

	for {

		message, ok := <-c.Message
		if !ok {
			log.Println(c.Message)
			return
		}

		c.Conn.WriteJSON(message)
	}
}

func (c *Client) ReadMessage(hub *Hub) {
	defer func() {
		hub.UnRegister <- c
		c.Conn.Close()
	}()

	for {
		_, m, err := c.Conn.ReadMessage()
		log.Println(string(m))
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error : %v", err)
			}
			break
		}

		msg := &Message{
			Content:  string(m),
			RoomID:   c.RoomID,
			Username: c.Username,
		}

		hub.Broadcast <- msg
	}

}
