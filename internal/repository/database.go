package database

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// NewDatabase - returns a pointer to a new database connection
func NewDatabase() (*gorm.DB, error) {
	//dbUsername := os.Getenv("DB_USER")
	//dbPassword := os.Getenv("DB_PASSWD")
	//dbHost := os.Getenv("DB_ADDR")
	//dbPort := os.Getenv("DB_PORT")
	//
	//// dbConnectionString := dbUsername + ":" + dbPassword + "@tcp(" + dbHost + ":" + strconv.Itoa(dbPort) + ")/" + dbTable
	//connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, dbUsername, dbPassword)
	//fmt.Println(connectionString)
	dbUser := os.Getenv("DB_USER")
	dbPasswd := os.Getenv("DB_PASSWD")
	dbAddr := os.Getenv("DB_ADDR")
	//dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbAddr, 5432, dbUser, dbPasswd, dbName)
	fmt.Println(psqlInfo)

	db, con_err := gorm.Open("postgres", psqlInfo)
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