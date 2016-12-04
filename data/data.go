package data

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

var Db *sql.DB

const (
	DB_HOST     = "localhost"
	DB_PORT     = "5433"
	DB_USER     = "postgres"
	DB_PASSWORD = ""
	DB_NAME     = "golangchina"
)

//init the connection to database using stdlib database/sql.
//When connect return with a error, the program will shutdown with a panic and log it.
func init() {
	var err error
	//sslmode set disable , we don't use ssl for local connection
	dbinfo := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)
	Db, err = sql.Open("postgres", dbinfo)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("SQL Connection Success(%s:%s)\n", DB_HOST, DB_PORT)
	}

	err = Db.Ping()
	if err != nil {
		log.Fatal("Error: Could not establish a connection with the database", err)
	}
	return
}
