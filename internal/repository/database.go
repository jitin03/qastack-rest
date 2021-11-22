package database

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// NewDatabase - returns a pointer to a new database connection
func NewDatabase() (*gorm.DB, error) {
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbTable := os.Getenv("DB_TABLE")
	dbPort := os.Getenv("DB_PORT")

	// dbConnectionString := dbUsername + ":" + dbPassword + "@tcp(" + dbHost + ":" + strconv.Itoa(dbPort) + ")/" + dbTable
	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, dbUsername, dbTable, dbPassword)
	fmt.Println(connectionString)

	db, con_err := gorm.Open("postgres", "host=localhost port=5433 user=postgres dbname=postgres sslmode=disable password=qastack")
	if con_err != nil {
		return db, con_err
	}

	if err := db.DB().Ping(); err != nil {
		return db, err
	}
	fmt.Println("Setting up new Database connection")
	return db, nil
}

// func CloseDb()(*gorm.DB)