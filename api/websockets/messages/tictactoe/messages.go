package tictactoe

import (
	"github.com/drew138/tictac/api/websockets/messages"
)

// Message is base message for tictactoe messaging
type Message struct {
	messages.Message
	UserID string `json:"userID"`
	RoomID string `json:"roomID"`
}
