package database

import (
	"fmt"
	"os"

	"gorm.io/gorm"
)

// DBConn database connection
var (
	DBConn *gorm.DB
)

// GetDSN format DSN string to be used in database related functions
func GetDSN() string {
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	DSN := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		host, port, username, name, password)
	return DSN // return data source name
}
