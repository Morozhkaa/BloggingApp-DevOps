package websocket

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	ID   string
	Name string
	Conn *websocket.Conn
	Pool *Pool
}

type Message struct {
	Type int    `json:"type"`
	Body string `json:"body"`
}

func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		messageType, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		if len(c.Pool.Buffer) < 200 && string(p) != "" {
			message := Message{Type: messageType, Body: string(p)}
			message.Body = fmt.Sprintf("%s:  %s", c.Name, message.Body)
			c.Pool.Broadcast <- message
			c.Pool.Buffer = append(c.Pool.Buffer, message)
			fmt.Printf("Message received : %+v\n", message)
			fmt.Printf("Length of buff: %+v\n", len(c.Pool.Buffer))
		}
	}
}
