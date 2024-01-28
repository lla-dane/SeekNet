package main

import (
	"sync"

	"github.com/fatih/color"

	"github.com/gorilla/websocket"
)

type Client struct {
	Conn     *websocket.Conn
	Username string
	msgColor *color.Color
}

var (
	clients = make(map[*websocket.Conn]*Client)
	mutex   sync.Mutex
)
