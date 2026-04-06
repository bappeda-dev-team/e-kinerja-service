package main

import (
	"database/sql"
	"log"
)

func SeedRoles(db *sql.DB) {
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
		log.Fatal("Gagal seed roles:", err)
	}

	log.Println("✅ Roles seeded successfully")
}
