package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/gorilla/websocket"
)

func (client *Client) broadcastMessage(message string) {
	mutex.Lock()
	defer mutex.Unlock()

	for _, allClient := range clients {

		if allClient.Conn == client.Conn {
			continue
		}

		err := allClient.Conn.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			log.Printf("Error broadcsting to %s: %s", client.Username, err)
			client.Conn.Close()
			delete(clients, client.Conn)
			continue
		}
	}
}

func (client *Client) getUsernamePassword() {

	paswd := white(bold("PASSWORD"))
	if err := client.Conn.WriteMessage(websocket.TextMessage, []byte(paswd)); err != nil {
		log.Println("Error sending the password prompt: ", err)
		return
	}

	_, paswdBytes, err := client.Conn.ReadMessage()
	if err != nil {
		log.Println(err)
		return
	}

	if strings.TrimSpace((string(paswdBytes))) != password {
		client.Conn.WriteMessage(websocket.TextMessage, []byte(red(bold("WRONG PASSWORD"))))
		log.Printf("FAILED LOGIN FROM %v", client.Conn.RemoteAddr())
		client.Conn.Close()
		return
	}

	username := white(bold("USERNAME"))
	if err := client.Conn.WriteMessage(websocket.TextMessage, []byte(username)); err != nil {
		log.Println("Error sending username promt: ", err)
		return
	}
	_, usernameBytes, err := client.Conn.ReadMessage()
	if err != nil {
		log.Println(err)
		return
	}
	client.Username = strings.TrimSpace(string(usernameBytes))
	mutex.Lock()
	clients[client.Conn] = client
	mutex.Unlock()

	msg := green(bold("CONNECTED TO THE FEED"))
	client.Conn.WriteMessage(websocket.TextMessage, []byte(msg))
}

func (client *Client) welcome() {
	welcomeMessage := fmt.Sprintf("[%s] has connected to the chat   -%v\n", client.Username, client.Conn.RemoteAddr())
	coloredWelcomeMessage := green(bold(welcomeMessage))
	client.broadcastMessage(coloredWelcomeMessage)
	log.Printf("A connection made from %v", client.Conn.RemoteAddr())
}

func (client *Client) handleMessages() {
	defer func() {
		exitMessage := (fmt.Sprintf("\n[%s] has disconnected %v\n", client.Username, client.Conn.RemoteAddr()))
		coloredExitMessage := red(exitMessage)
		client.Conn.Close()
		mutex.Lock()
		delete(clients, client.Conn)
		mutex.Unlock()
		client.broadcastMessage(bold(coloredExitMessage))
	}()

	for {
		_, msg, err := client.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}

		trimmedMsg := strings.TrimSpace(string(msg))

		if trimmedMsg != "" {
			exMessage := fmt.Sprintf("[%s]: %s", client.Username, msg)
			coloredExMessage := client.msgColor.Sprintf(exMessage)
			log.Println(bold(coloredExMessage))
			client.broadcastMessage(bold(coloredExMessage))
		}

	}

}