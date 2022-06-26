package main

import (
	"fmt"
	"net/http"

	"websocket_sample/pkg/websocket"
)

func serveWebsocket(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}

	client := &websocket.Client{
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client
	client.Read()
}

func configRoutes() {

	pool := websocket.Pool{
		Register:   make(chan *websocket.Client),
		Unregister: make(chan *websocket.Client),
		Clients:    make(map[*websocket.Client]bool),
		Broadcast:  make(chan websocket.Message),
	}

	go pool.Start()

	http.HandleFunc("/websocket", func(w http.ResponseWriter, r *http.Request) {
		serveWebsocket(&pool, w, r)
	})
}

func main() {
	fmt.Println("Distributed App")
	configRoutes()
	http.ListenAndServe(":8080", nil)
}
