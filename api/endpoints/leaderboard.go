package endpoints

import (
	"encoding/json"
	"net/http"

	"github.com/drew138/tictac/api/status"
	"github.com/drew138/tictac/database"
	"github.com/drew138/tictac/database/models"
)

// struct of users required fields
type userRank struct {
	Name string
	Wins int
}

// GetLeaderboard - retrieve names of players with highest scores
func GetLeaderboard(w http.ResponseWriter, r *http.Request) {
	var users []userRank

	err := database.DBConn.Model(&models.User{}).Order("Wins desc").Limit(100).Find(users).Error
	if err != nil {
		status.RespondStatus(w, 500, err)
		return
	}
	json.NewEncoder(w).Encode(users)
}
