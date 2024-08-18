package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// Upgrades connection from HTTP to WS
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	// Check origin of the connection to process requests from react client
	CheckOrigin: func(r *http.Request) bool { return true },
}

// Function responsible for listening for messages over connection
func reader(conn *websocket.Conn) {
	for {
		// Read/error check message
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		// print bytes (message) as string
		fmt.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println((err))
			return
		}
	}
}

// WS endpoint
func serveWebSocket(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host)

	// Upgrade request to WS
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	// Read messages if creating WS was success
	reader(ws)
}

func setup() {
	http.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Test")
		})

	// map /ws request to function to serve websocket
	http.HandleFunc("/ws", serveWebSocket)
}

func main() {
	fmt.Println("Message App")
	setup()
	http.ListenAndServe(":8080", nil)
}
