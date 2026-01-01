package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

// InitDB initializes driver and connection pool for connecting to postgres
// database with connection information
func InitDB() *sql.DB {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Failed to ping DB:", err)
	}

	return db
}
