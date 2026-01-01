package migrations_test

import (
	"backend/internal/db"
	"backend/internal/migrations"
	"database/sql"
	"fmt"
	"log"
	"testing"
)

// Helper to initialize database
func setupTestDB(t *testing.T) *sql.DB {
	db, err := db.InitDB()
	if err != nil {
		log.Println("DB connection failed:", err)
	}
	const dropTableSQL = `
	DROP TABLE IF EXISTS 
			member, 
			staff, 
			message, 
			talk_user, 
			talk_user_member,
    	template, 
			fanletter, 
			blog, 
			notification, 
			official_news, 
			message_read, 
			schema_migrations 
	CASCADE;
	`
	if _, err := db.Exec(dropTableSQL); err != nil {
		t.Fatal("Failed to truncate tables:", err)
	}

	return db
}

// Test migrations
func TestMigrationsTablesExist(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	if err := migrations.RunMigrations(db); err != nil {
		log.Fatalf("Failed to apply migration: %v", err)
	}

	expectedTables := []string{
		"member",
		"staff",
		"message",
		"talk_user",
		"talk_user_member",
		"template",
		"fanletter",
		"blog",
		"notification",
		"official_news",
		"message_read",
		"schema_migrations",
	}

	var missing []string

	// Check that all tables are generated
	for _, table := range expectedTables {
		var exists bool
		query := fmt.Sprintf(`
		SELECT EXISTS (
				SELECT 1
				FROM information_schema.tables
				WHERE table_schema = 'public'
					AND table_name = '%s'
		);
		`, table)

		err := db.QueryRow(query).Scan(&exists)
		if err != nil {
			t.Fatalf("failed to query table %s: %v", table, err)
		}

		if !exists {
			missing = append(missing, table)
		}
	}

	if len(missing) > 0 {
		t.Errorf("The following expected tables do not exist: %v", missing)
	} else {
		t.Log("All expected tables exist.")
	}
}
