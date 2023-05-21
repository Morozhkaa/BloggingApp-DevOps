package websocket

import (
	"fmt"
)

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan Message
	Buffer     []Message
}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan Message),
		Buffer:     []Message{},
	}
}

func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true
			notification := fmt.Sprintf("New User '%s' Joined", client.Name)
			fmt.Println("size of connection pool:", len(pool.Clients))
			for client := range pool.Clients {
				fmt.Println(client)
				client.Conn.WriteJSON(Message{Type: 1, Body: notification})
			}
		case client := <-pool.Unregister:
			delete(pool.Clients, client)
			notification := fmt.Sprintf("User '%s' Disconnected", client.Name)
			fmt.Println("size of connection pool:", len(pool.Clients))
			for client := range pool.Clients {
				client.Conn.WriteJSON(Message{Type: 1, Body: notification})
			}
		case message := <-pool.Broadcast:
			fmt.Println("Sending message to all clients in the pool")
			for client := range pool.Clients {
				fmt.Println(client)
				if err := client.Conn.WriteJSON(message); err != nil {
					fmt.Println(err)
					return
				}
			}
		}

	}
}
