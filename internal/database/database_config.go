package database

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

var (
	db *sql.DB
)

type Credentials struct {
	User     string
	Password string
}

func Connect() {

	log.Infoln("Starting to connect to database")

	d, err := sql.Open("mysql", os.Getenv("MYSQL_KUBE"))

	if err != nil {
		log.Fatalf("Error connecting to database: %s", err.Error())
	} else {
		log.Infoln("Connected to database")
	}

	db = d
}

func DB() *sql.DB {

	if db == nil {
		Connect()
	}
	return db
}
