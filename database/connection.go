package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	// requirement for DBConn
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConnectDB initialize connection to postgres database
func ConnectDB() {
	var err error
	DBConn, err = gorm.Open(postgres.Open(GetDSN()), &gorm.Config{})
	if err != nil {
		panic("Failed to establish connection to database")

	}
	fmt.Println("Established connection succesfully to database")
	// DBConn
}
