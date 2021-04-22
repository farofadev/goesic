package lib

import (
	"database/sql"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var defaultDBConnection *sql.DB

func DBConnect() *sql.DB {
	config := mysql.NewConfig()
	config.Net = "tcp"
	config.Addr = os.Getenv("DB_HOST")
	config.DBName = os.Getenv("DB_NAME")
	config.User = os.Getenv("DB_USER")
	config.Passwd = os.Getenv("DB_PASSWORD")

	database, err := sql.Open("mysql", config.FormatDSN())

	if (err != nil) {
		log.Fatal("Error when trying to connect to database", err)
	}

	return database
}

func SetDefaultDBConnection(db *sql.DB) *sql.DB {
	defaultDBConnection = db

	return db
}

func DBConnectAndSetDefault() *sql.DB {
	return SetDefaultDBConnection(DBConnect())
}

func GetDefaultDBConnection() *sql.DB {
	return defaultDBConnection
}