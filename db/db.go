package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *sql.DB
var err error

func Connection() {

	dialect := os.Getenv("DIALECT")
	host := os.Getenv("HOST")
	dbPort := os.Getenv("DBPORT")
	user := os.Getenv("USER")
	dbname := os.Getenv("NAME")
	dbpassword := os.Getenv("PASSWORD")

	// Database connection string
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, user, dbname, dbpassword, dbPort)

	DB, err = sql.Open(dialect, dbURI)

	if err != nil {
		panic("Conexão mal sucedida  :(")
	} else {
		fmt.Println("Conexão ok")
	}
}