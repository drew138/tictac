package websockets

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/drew138/tictac/api/websockets/connections"

	messages "github.com/drew138/tictac/api/websockets/messages/tictactoe"
	"github.com/gorilla/websocket"
)

// Upgrader -
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
	go handleMessages(conn, c)
}

func handleMessages(conn *websocket.Conn, c *connections.Connections) error {
	log.Println("Websocket connection established.")
	defer func() {
		if err := conn.Close(); err != nil {
			log.Fatalln("Failed to close websocket connection: ", err.Error())
		} else {
			log.Println("Websocket connection closed.")
		}
	}()
	for {
		var newMessage messages.Message
		if err := conn.ReadJSON(&newMessage); err != nil {
			log.Fatalln("Error when reading websocket message: ", err.Error())
			return err
		}
		j, _ := json.Marshal(newMessage)
		log.Println("Message Recieved: ", string(j))
		switch newMessage.Action {
		case "connect":
			c.Connect <- newMessage.UserID
			break
		case "disconnect":
			c.Disconnect <- newMessage.UserID
			break
		default:
			log.Println("Unknown message recieved: ", string(j))
			break
		}
	}
}

// https://tutorialedge.net/projects/chat-system-in-go-and-react/
