package main

import (
	"database/sql"
	"log"
)

func SeedMasterPemda(db *sql.DB) {
	query := `
	INSERT INTO master_pemda (name)
	VALUES
	('Pemda tangerang'),
	('Pemda Sukoharjo'),
	('Pemda Semarang')
	ON CONFLICT (name) DO NOTHING;
	`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal("Gagal seed master pemda:", err)
	}

	log.Println("✅ Master Pemda seeded successfully")
}