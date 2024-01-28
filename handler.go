package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// Implement origin checking logic here
		return true
	},
}

func (client *Client) init() {

	intro := yellow(bold("\n----------\n     WELCOME TO THE FEED\n     HERE YOU TALK IN THE TERMINAL\n     ENTER THE PASSWORD AND USERNAME TO START\n----------\n"))

	if err := client.Conn.WriteMessage(websocket.TextMessage, []byte(intro)); err != nil {
		log.Fatal(" Error connecting: ", err)
	}
}

func wsHandler(w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(colors))
	randomColor := colors[randomIndex]

	client := &Client{
		Conn:     conn,
		msgColor: randomColor,
	}

	client.init()

	client.getUsernamePassword()
	client.welcome()
	client.handleMessages()

}
