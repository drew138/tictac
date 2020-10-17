package dependencies

import (
	"github.com/drew138/tictac/api/websockets/connections"
)

// Dependencies holds all dependencies required by project
type Dependencies struct {
	WebsocketConnections *connections.Connections
}
