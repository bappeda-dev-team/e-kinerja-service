package main

import (
	"database/sql"
	"log"
)

func SeedMasterAplikasi(db *sql.DB) {
	query := `
	INSERT INTO master_aplikasi (name)
	VALUES 
	('aplikasi tracer study'),
	('sistem sensor suhu'),
	('website monitoring');
	`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal("Gagal seed master aplikasi:", err)
	}

	log.Println("✅ Master Aplikasi seeded successfully")
}