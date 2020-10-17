package connections

import (
	"log"
)

// Connections holds channels for interacting with connection pool of users
type Connections struct {
	Connect        chan string
	Disconnect     chan string
	connectedUsers map[string]bool
}

// StartConnectionTracking starts connection workers and returns Connections struct
func StartConnectionTracking() *Connections {
	connectionPool := Connections{
		Connect:        make(chan string),
		Disconnect:     make(chan string),
		connectedUsers: make(map[string]bool),
	}
	go connectionPool.startConnectionsWorker()
	return &connectionPool
}

func (c *Connections) startConnectionsWorker() {
	for {
		select {
		case id := <-c.Connect:
			c.connectedUsers[id] = true
			log.Println("Added user", id, "to connection pool")
			break
		case id := <-c.Disconnect:
			delete(c.connectedUsers, id)
			log.Println("Removed user", id, "from connection pool")
			break
		default:
			log.Println("Unrecognized message sent to connection worker")
			break
		}
	}
}
