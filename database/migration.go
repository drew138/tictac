package database

import (
	"fmt"

	"github.com/drew138/tictac/database/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// AutoMigrateDB initialize connection to postgres database
func AutoMigrateDB() {
	var err error
	DBConn, err = gorm.Open(postgres.Open(GetDSN()), &gorm.Config{})
	if err != nil {
		panic("Failed to establish connection to database")
	}
	fmt.Println("Established connection succesfully to database")
	DBConn.AutoMigrate(&models.User{}, &models.Game{}) // TODO add remaining models
	fmt.Println("Database Migrated Succesfully")
	// DBConn
}
