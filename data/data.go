package data

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	_ "utils"
)

var Db *sql.DB

//InitDataBaseConnect: init the connection to database using stdlib database/sql.
//When connect return with a error, the program will shutdown with a panic and log it.
func InitDataBaseConnect() {
	var err error
	var env string = viper.GetString("env")
	var databaseInfo map[string]string = viper.GetStringMapString(env + ".database.postgres")
	DB_HOST := databaseInfo["db_host"]
	DB_PORT := databaseInfo["db_port"]
	DB_USER := databaseInfo["db_user"]
	DB_PASSWORD := databaseInfo["db_password"]
	DB_NAME := databaseInfo["db_name"]

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
