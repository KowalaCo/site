package api

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

//ConnectDatabase postgresql
func ConnectDatabase() {
	var err error
	db, err = sql.Open("postgres", Config.Db)
	if err != nil {
		log.Print("Postgresql failed to connect!")
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Print("Postgresql failed to connect!")
		log.Fatal(err)
	}
}
