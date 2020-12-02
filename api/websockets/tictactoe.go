package websockets

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/drew138/tictac/api/websockets/connections"
	messages "github.com/drew138/tictac/api/websockets/messages/tictactoe"
	"github.com/gorilla/websocket"
	"github.com/segmentio/ksuid"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

// Upgrader
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

// HandleConnection handles initial websocket connection
func HandleConnection(w http.ResponseWriter, r *http.Request, c *connections.Connections) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("Error upgrading websocket connection: ", err.Error())
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(&map[string]string{"Error": err.Error()})
	}
	user := &connections.ConnectedUser{
		UserID: ksuid.New().String(),
		Conn:   conn,
	}
	c.Connect <- user
	go handleMessages(user, c)
}

func handleMessages(user *connections.ConnectedUser, c *connections.Connections) {
	log.Println("Websocket connection established.")
	// userID := user.UserID
	conn := user.Conn
	conn.SetReadLimit(maxMessageSize)
	conn.SetReadDeadline(time.Now().Add(pongWait))
	conn.SetPongHandler(func(string) error { conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	defer func() {
		if err := conn.Close(); err != nil {
			log.Fatalln("Failed to close websocket connection: ", err.Error())
		} else {
			c.Disconnect <- user
			log.Println("Websocket connection closed.")
		}
	}()
	for {
		var newMessage messages.Message
		if err := conn.ReadJSON(&newMessage); err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Unexpected websocket close: %v", err)
			}
			break
		}
		j, _ := json.Marshal(newMessage)
		log.Println("Message Recieved: ", string(j))
		switch newMessage.Action {
		default:
			log.Println("Unknown message recieved: ", string(j))
			break
		}
	}
}

// https://tutorialedge.net/projects/chat-system-in-go-and-react/
