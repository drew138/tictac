package websockets

import (
	"encoding/json"
	"fmt"
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
	pongWait = 10 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
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
		UserID:    ksuid.New().String(),
		Conn:      conn,
		SendQueue: make(chan *messages.Message),
	}
	c.Connect <- user
	go handleMessages(user, c)
	go sendMessagesWorker(user, c)
	log.Println("Websocket connection established.")
}

func handleMessages(user *connections.ConnectedUser, c *connections.Connections) {
	userID := user.UserID
	conn := user.Conn
	conn.SetReadLimit(maxMessageSize)
	// TODO: Investigate ping/pong system
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
			log.Printf("Unexpected message format: %v", err)
			break
		}
		j, _ := json.Marshal(newMessage)
		log.Println("Message Recieved: ", string(j))
		switch newMessage.Action {
		case "privateMessage", "gameAction":
			otherUser := c.GetConnectedUser(newMessage.RecipientID)
			if otherUser != nil {
				outMsg := &messages.Message{
					Action:   newMessage.Action,
					SenderID: userID,
					Body:     newMessage.Body,
				}
				otherUser.SendQueue <- outMsg
				user.SendQueue <- &newMessage
			} else {
				outMsg := &messages.Message{
					Action: "error",
					Body:   "User is not connected",
				}
				user.SendQueue <- outMsg
			}
		default:
			log.Println("Unknown message recieved: ", string(j))
			break
		}
	}
}

func sendMessagesWorker(user *connections.ConnectedUser, c *connections.Connections) {
	conn := user.Conn
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		conn.Close()
	}()
	for {
		select {
		case message, ok := <-user.SendQueue:
			conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			messageJSON, err := json.Marshal(*message)
			if err != nil {
				log.Printf("Error parsing JSON: %v", err)
			}
			w.Write(messageJSON)

			// Add queued chat messages to the current websocket message.
			n := len(user.SendQueue)
			for i := 0; i < n; i++ {
				w.Write(newline)
				newMessage, err := json.Marshal(<-user.SendQueue)
				if err != nil {
					log.Printf("Error parsing JSON: %v", err)
				}
				w.Write(newMessage)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}

// https://tutorialedge.net/projects/chat-system-in-go-and-react/
