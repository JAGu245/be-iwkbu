package config

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite", "iwkbu.db")
	if err != nil {
		log.Fatal(err)
	}

	// Create table if not exists
	schema := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT UNIQUE NOT NULL,
		fullname TEXT DEFAULT '',
		password TEXT NOT NULL,
		role TEXT DEFAULT 'user',
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	_, err = DB.Exec(schema)
	if err != nil {
		log.Fatal(err)
	}

	// Add fullname column if it doesn't exist (proactive migration)
	_, _ = DB.Exec("ALTER TABLE users ADD COLUMN fullname TEXT DEFAULT ''")

	log.Println("Database initialized successfully")
}
