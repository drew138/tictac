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
	DSN := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s",
		host, username, name, password)
	return DSN // return data source name
}
