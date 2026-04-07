package main

import (
	"database/sql"
	"log"
)

func SeedMasterAplikasi(db *sql.DB) {
	query := `
	INSERT INTO master_aplikasi (name, link)
	SELECT name, link FROM (VALUES
		('Mahakam Hulu', 'https://kk.kertaskerja.cc/'),
		('Kab Madiun', 'https://kab.kertaskerja.cc/'),
		('Semarang', 'https://semarang.kertaskerja.cc/'),
		('Sleman', 'https://kabsleman.kertaskerja.cc/'),
		('Gresik', 'https://kabgresik.kertaskerja.cc'),
		('Sragen', 'https://kabsragen.kertaskerja.cc/'),
		('Kota Pasuruan', 'https://kotapasuruan.kertaskerja.cc/'),
		('Bontang', 'https://kotabontang.kertaskerja.cc'),
		('Ponorogo', 'https://kabponorogo.kertaskerja.cc/login'),
		('Kab Sukoharjo', 'https://kabsukoharjo.kertaskerja.cc'),
		('Kota Probolinggo', 'https://kotaprobolinggo.kertaskerja.cc')
	) AS v(name, link)
	WHERE NOT EXISTS (
		SELECT 1 FROM master_aplikasi WHERE master_aplikasi.name = v.name
	);
	`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal("Gagal seed master aplikasi:", err)
	}

	log.Println("✅ Master Aplikasi seeded successfully")
}
