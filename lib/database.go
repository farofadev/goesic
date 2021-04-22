package lib

import (
	"database/sql"
	"log"
	"os"
	"time"

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

	db, err := sql.Open("mysql", config.FormatDSN())

	if (err != nil) {
		log.Fatal("Error when trying to connect to database", err)
	}

	db.SetMaxIdleConns(64)
	db.SetMaxOpenConns(64)
	db.SetConnMaxLifetime(-1)

	return db
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

func LoopVerifyDefaultDBConnection() {
	for {
		db := GetDefaultDBConnection();

		log.Println("Verificando banco de dados")

		_, err := db.Query("SHOW TABLES;")

		if err != nil {
			log.Println("Conexão falhou, tentando reconectar...")

			DBConnectAndSetDefault();
		}

		time.Sleep(30 * time.Second)
	}
}