package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dbUser = "jlum_user"
	dbName = "jlum_db"
	dbHost = "127.0.0.1"
	dbPort = "3306"
)

var DB *sql.DB

func InitDB(dbPassword string) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		dbUser, dbPassword, dbHost, dbPort, dbName)
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to MySQL with DSN '%s': %v", dsn, err)
	}
	if err := DB.Ping(); err != nil {
		log.Fatalf("DB ping failure: %v", err)
	}
	log.Println("MySQL connnected successfully")
}
