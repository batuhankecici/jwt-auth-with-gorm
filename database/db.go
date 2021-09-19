package database

import (
	"fmt"
	"jwt-auth-with-gorm/models"
	"log"
	"os"

	"github.com/jinzhu/gorm"
)

func ConnectDB() (*gorm.DB, error) {

	dialect := os.Getenv("DIALECT")
	host := os.Getenv("HOST")
	dbPort := os.Getenv("DBPORT")
	user := os.Getenv("USER")
	dbName := os.Getenv("NAME")
	password := os.Getenv("PASSWORD")

	// Db ConnectionString
	connectionString := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, user, dbName, password, dbPort)
	db, err := gorm.Open(dialect, connectionString)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func AutoMigrate() (*gorm.DB, error) {
	conn, err := ConnectDB()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	conn.AutoMigrate(models.Book{})
	log.Println("Migration has been processed")
	return conn, nil
}
