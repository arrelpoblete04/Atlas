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

func Connect(credentials Credentials) {

	log.Infoln("Starting to connect to database")
	log.Infof("MySQL Kube: %s", os.Getenv("MYSQL_KUBE"))
	log.Infof("DB Username: %s", credentials.User)
	log.Infof("DB Password: %s", credentials.Password)

	d, err := sql.Open("mysql", os.Getenv("MYSQL_KUBE"))

	if err != nil {
		log.Fatalf("Error connecting to dasssstabase: %s", err.Error())
	} else {
		log.Infoln("Connected to database")
	}

	db = d
}

func DB(username, password string) *sql.DB {

	credentials := Credentials{
		User:     username,
		Password: password,
	}

	if db == nil {
		Connect(credentials)
	}
	return db
}
