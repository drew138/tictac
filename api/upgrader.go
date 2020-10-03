package api

import "github.com/gorilla/websocket"

// Upgrader -
var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// https://tutorialedge.net/projects/chat-system-in-go-and-react/
