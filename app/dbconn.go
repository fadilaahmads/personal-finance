package app

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func DBConn() *sql.DB {
	var (
		host     = os.Getenv("DB_HOST")
		port     = os.Getenv("DB_PORT")
		username = os.Getenv("DB_USERNAME")
		password = os.Getenv("DB_PASSWORD")
		dbname   = os.Getenv("DB_NAME")
	)
	dbinfo := fmt.Sprintf("%s:%s@(%s:%s)/%s?parsetime?=true", username, password, host, port, dbname)
	db, err := sql.Open("mysql", dbinfo)
	if err != nil {
		log.Panic(err)
	}
	if err := db.Ping(); err != nil {
		log.Panic(err)
	}
	return db
}
