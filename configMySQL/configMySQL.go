package configmysql

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	username = "root"
	password = "P@ssw0rd"
	hostname = "localhost:3306"
	dbname   = "ms_account_dev"
)

// Connect db
func Connect() *sql.DB {
	db, err := sql.Open("mysql", confMysql(dbname))

	if err != nil {
		log.Fatal(err)
	}

	return db
}

// GetValue ...
func GetValue() string {
	return "Hello from this another package"
}

func confMysql(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
}
