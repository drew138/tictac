package websockets

import (
	"encoding/json"
	"log"
	"net/http"

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
func HandleConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("Error upgrading websocket connection: ", err.Error())
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(&map[string]string{"Error": err.Error()})
	}
	go handleMessages(conn)
}

func handleMessages(conn *websocket.Conn) error {
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
	}
}

// https://tutorialedge.net/projects/chat-system-in-go-and-react/
