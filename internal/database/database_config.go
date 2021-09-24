package database

import (
	// "database/sql"
	// "database/sql"

	"database/sql"

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

	d, err := sql.Open("mysql", "root:password@tcp(10.1.0.57:3306/colors")

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
