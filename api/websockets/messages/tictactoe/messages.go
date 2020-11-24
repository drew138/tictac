package tictactoe

// Message is base message for tictactoe messaging
type Message struct {
	Action      string `json:"action"`
	UserID      string `json:"userID"`
	SenderID    string `json:"senderID"`
	RecipientID string `json:"recipientID"`
	RoomID      string `json:"roomID"`
	Body        string `json:"body"`
}
