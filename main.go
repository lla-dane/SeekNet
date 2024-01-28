package main

import (
	"fmt"
	"log"
	"net/http"
)

var password string

func main() {

	fmt.Println(green(bold("CREATE PASSWORD FOR THE ROOM:")))
	fmt.Scan(&password)

	http.HandleFunc("/ws", wsHandler)
	log.Println("SERVER RUNNING ON PORT 5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
