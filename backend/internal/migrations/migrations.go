package migrations

import (
	"database/sql"
	"log"
	"os"
	"strings"
)

// RunMigrations runs all migration code in migrations directory that has not
// been applied yet.
func RunMigrations(db *sql.DB) error {
	// Ensure schema_migrations table exists
	const schemaMigrationSQL = `
	CREATE TABLE IF NOT EXISTS schema_migrations (
			version TEXT PRIMARY KEY,
			applied_at TIMESTAMP NOT NULL DEFAULT NOW()
	);
	`
	if _, err := db.Exec(schemaMigrationSQL); err != nil {
		return err
	}

	// Get applied migrations
	applied := make(map[string]bool)
	rows, err := db.Query(
		`SELECT version FROM schema_migrations`,
	)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var version string
		if err := rows.Scan(&version); err != nil {
			return err
		}
		applied[version] = true
	}

	// Read migration files
	files, err := os.ReadDir(".")
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() || !strings.HasSuffix(file.Name(), ".sql") {
			continue
		}
		name := file.Name()
		version := strings.SplitN(name, "_", 2)[0]

		if applied[version] {
			continue
		}

		data, err := os.ReadFile(name)
		if err != nil {
			return err
		}

		// extract +up section
		parts := strings.Split(string(data), "--- +up")
		if len(parts) < 2 {
			log.Fatalf("Migration %s missing +up section", name)
		}
		upSQL := strings.Split(parts[1], "--- +down")[0]
		upSQL = strings.TrimSpace(upSQL)

		if err := applyMigrations(db, version, upSQL); err != nil {
			return err
		}

		log.Println("Applied migration:", name)
	}
	return nil
}

// applyMigrations is a helper for appling singe migration to database
func applyMigrations(db *sql.DB, version string, upSQL string) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	if _, err := tx.Exec(upSQL); err != nil {
		tx.Rollback()
		return err
	}

	if _, err := tx.Exec(
		`INSERT INTO schema_migrations (version) VALUES ($1)`,
		version,
	); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
