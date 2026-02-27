package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:psg2026@127.0.0.1:5432/db_internal?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	seedRoles(db)
}

func seedRoles(db *sql.DB) {
	query := `
	INSERT INTO roles (name, description)
	VALUES 
	('super_admin', 'Super Administrator'),
	('admin', 'Administrator'),
	('programmer', 'Programmer'),
	('level2', 'Programmer & Verifikator')
	ON CONFLICT (name) DO NOTHING;
	`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Roles seeded successfully")
}



