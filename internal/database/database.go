package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
)
// NewDatabase Connection Func return pointer to DB connection
func NewDatabase() (*gorm.DB, error) {
	fmt.Println("Setting up New DB connection")

	dbUsername := os.getenv("DB_USERNAME")
	dbPassword := os.getenv("DB_PASSWORD")
	dbHost := os.getenv("DB_HOST")
	dbTable := os.getenv("DB_TABLE")
	dbPort := os.getenv("DB_PORT")

	connectString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUsername, dbTable, dbPassword)

	db, err := gorm.Open("postgres", connectString)
	if err != nil {
		return db, err
	}

	if err := db.DB().Ping(); err != nil {
		return db,err
	}

	return db, nil
}