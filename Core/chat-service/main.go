package main

import (
	"chat/websocket"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/cloudflare/cfssl/log"
)

func random(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	if min > max {
		return min
	} else {
		return rand.Intn(max-min) + min
	}
}

var names_arr = [...]string{
	"Меркурий",
	"Венера",
	"Земля",
	"Марс",
	"Юпитер",
	"Сатурн",
	"Уран",
	"Нептун",
	"Солнце",
}

var names_count = map[string]int{
	"Меркурий": 0,
	"Венера":   0,
	"Земля":    0,
	"Марс":     0,
	"Юпитер":   0,
	"Сатурн":   0,
	"Уран":     0,
	"Нептун":   0,
}

func serveChat(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println("websocket endpoint reached")

	conn, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}

	name := names_arr[random(0, 8)]
	names_count[name] += 1
	client := &websocket.Client{
		Name: fmt.Sprintf("%s-%d", name, names_count[name]),
		Conn: conn,
		Pool: pool,
	}
	pool.Register <- client
	client.Read()
}

func setupRoutes() {
	pool := websocket.NewPool()
	go pool.Start()

	http.HandleFunc("/api/chat-service/v1/chat", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		serveChat(pool, w, r)
	})
}

func main() {
	log.Info("Start chat service")
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}
