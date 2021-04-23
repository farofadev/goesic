package lib

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
)

var defaultDBConnection *sql.DB

func DBConnect() (*sql.DB, error) {
	config := mysql.NewConfig()
	config.Net = "tcp"
	config.Addr = os.Getenv("DB_HOST")
	config.DBName = os.Getenv("DB_NAME")
	config.User = os.Getenv("DB_USER")
	config.Passwd = os.Getenv("DB_PASSWORD")

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

func SetDefaultDBConnection(db *sql.DB) *sql.DB {
	defaultDBConnection = db

	return db
}

func DBConnectAndSetDefault() (*sql.DB, error) {
	db, err := DBConnect()

	return SetDefaultDBConnection(db), err
}

func GetDefaultDBConnection() *sql.DB {
	return defaultDBConnection
}

func LoopVerifyDefaultDBConnection() error {
	previous := false

	for {
		db := GetDefaultDBConnection()
		_, err := db.Query("SHOW TABLES;")

		if err != nil {
			log.Println("A conexão com o Banco de Dados falhou: ", err)

			connection, cerr := DBConnectAndSetDefault()

			if cerr != nil|| connection == nil {
				previous = true
				log.Println("Falhou ao reconectar...", cerr)
			} else {
				previous = false
			}
		} 

		if (previous) {
			time.Sleep(5 * time.Second)
		} else {
			time.Sleep(30 * time.Second)
		}
	}
}