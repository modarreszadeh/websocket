package websocket

import "fmt"

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan Message
}

func (pool *Pool) Start() {
	for {
		select {

		case client := <-pool.Register:
			pool.Clients[client] = true
			fmt.Println("Size of Connection Pool: ", len(pool.Clients))

		case client := <-pool.Unregister:
			delete(pool.Clients, client)
			fmt.Println("Size of Connection Pool: ", len(pool.Clients))

		case message := <-pool.Broadcast:
			fmt.Println("Sending message to all clients in Pool")

			for client := range pool.Clients {
				if err := client.Conn.WriteJSON(message.Body + client.ID); err != nil {
					fmt.Println(err)
					return
				}
			}
		}
	}
}
