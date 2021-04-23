package database

import (
	"database/sql"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

func GetDefaultDBConfig() *mysql.Config {
	config := mysql.NewConfig()
	config.Net = "tcp"
	config.Addr = os.Getenv("DB_HOST")
	config.DBName = os.Getenv("DB_NAME")
	config.User = os.Getenv("DB_USER")
	config.Passwd = os.Getenv("DB_PASSWORD")

	return config
}

func DBConnectDefault() (*sql.DB, error){
	return DBConnect(GetDefaultDBConfig())
}

func DBConnect(config *mysql.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", config.FormatDSN())

	if (err != nil) {
		log.Println("Error when trying to connect to database", err)
		return db, err
	}

	db.SetMaxIdleConns(64)
	db.SetMaxOpenConns(64)
	db.SetConnMaxLifetime(-1)

	return db, err
}
