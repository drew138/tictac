package websockets

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// Upgrader -
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
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

func handleMessages(conn *websocket.Conn) {

}

// https://tutorialedge.net/projects/chat-system-in-go-and-react/
