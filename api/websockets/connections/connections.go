package connections

import (
	"log"

	messages "github.com/drew138/tictac/api/websockets/messages/tictactoe"
	"github.com/gorilla/websocket"
)

// Connections holds channels for interacting with connection pool of users
type Connections struct {
	Connect        chan *ConnectedUser
	Disconnect     chan *ConnectedUser
	connectedUsers map[string]*ConnectedUser
}

// ConnectedUser holds information of a connected user, as well as their websocket connection
type ConnectedUser struct {
	UserID    string
	Conn      *websocket.Conn
	SendQueue chan *messages.Message
}

// StartConnectionTracking starts connection workers and returns Connections struct
func StartConnectionTracking() *Connections {
	connectionPool := Connections{
		Connect:        make(chan *ConnectedUser),
		Disconnect:     make(chan *ConnectedUser),
		connectedUsers: make(map[string]*ConnectedUser),
	}
	go connectionPool.startConnectionsWorker()
	return &connectionPool
}

func (c *Connections) startConnectionsWorker() {
	for {
		select {
		case user := <-c.Connect:
			c.connectedUsers[user.UserID] = user
			log.Println("Added user", user.UserID, "to connection pool")
			break
		case user := <-c.Disconnect:
			delete(c.connectedUsers, user.UserID)
			close(user.SendQueue)
			log.Println("Removed user", user.UserID, "from connection pool")
			break
		}
	}
}

// GetConnectedUser returns the conected user
func (c *Connections) GetConnectedUser(id string) *ConnectedUser {
	return c.connectedUsers[id]
}
