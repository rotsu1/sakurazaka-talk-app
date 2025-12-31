package db_test

import (
	"backend/internal/db"
	"testing"
)

// TestInitDB tests database initialization
func TestInitDB(t *testing.T) {
	db, err := db.InitDB()
	if err != nil {
		t.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		t.Fatalf("Failed to ping database: %v", err)
	}
}
