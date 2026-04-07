package main

import (
	"database/sql"
	"log"
)

// SeedWorkflow menanam data lengkap: permintaan → distribusi → pelaksana → laporan → verifikasi
// Bergantung pada: SeedRoles, SeedUsers, SeedMasterPemda, SeedMasterAplikasi
func SeedWorkflow(db *sql.DB) {
	// Ambil ID yang diperlukan
	superAdminID := mustGetUserID(db, "superadmin")
	adminID := mustGetUserID(db, "admin")
	programmerID := mustGetUserID(db, "programmer")
	verifikatorID := mustGetUserID(db, "verifikator")
	pemdaID := mustGetFirstID(db, "master_pemda")
	aplikasiID := mustGetFirstID(db, "master_aplikasi")

	// 1. Permintaan
	var permintaanID string
	err := db.QueryRow(`
		INSERT INTO permintaan (pemda_id, aplikasi_id, menu, kondisi_awal, kondisi_diharapkan,
		                        tanggal_pesanan, tanggal_deadline, created_by)
		VALUES ($1, $2, 'Dashboard Utama', 'Belum ada fitur export data',
		        'Tersedia tombol export ke Excel dan PDF',
		        '2026-03-01', '2026-04-01', $3)
		RETURNING id
	`, pemdaID, aplikasiID, superAdminID).Scan(&permintaanID)
	if err != nil {
		log.Fatal("Gagal seed permintaan:", err)
	}
	log.Println("✅ Permintaan seeded, id:", permintaanID)

	// 2. Distribusi (oleh admin)
	var distribusiID string
	err = db.QueryRow(`
		INSERT INTO distribusi (permintaan_id, admin_id, komentar)
		VALUES ($1, $2, 'Didistribusikan ke programmer terpilih')
		RETURNING id
	`, permintaanID, adminID).Scan(&distribusiID)
	if err != nil {
		log.Fatal("Gagal seed distribusi:", err)
	}
	log.Println("✅ Distribusi seeded, id:", distribusiID)

	// 3. Pelaksana (programmer ditugaskan)
	_, err = db.Exec(`
		INSERT INTO distribusi_pelaksana (distribusi_id, programmer_id)
		VALUES ($1, $2)
		ON CONFLICT (distribusi_id, programmer_id) DO NOTHING
	`, distribusiID, programmerID)
	if err != nil {
		log.Fatal("Gagal seed pelaksana:", err)
	}
	log.Println("✅ Pelaksana seeded (programmer:", programmerID, ")")

	// 4. Laporan kinerja (oleh programmer)
	var laporanID string
	err = db.QueryRow(`
		INSERT INTO laporan_kinerja (permintaan_id, programmer_id, laporan_progress, status)
		VALUES ($1, $2, 'Sudah selesai implementasi fitur export Excel dan PDF', 'hijau')
		RETURNING id
	`, permintaanID, programmerID).Scan(&laporanID)
	if err != nil {
		log.Fatal("Gagal seed laporan:", err)
	}
	log.Println("✅ Laporan seeded, id:", laporanID)

	// 5. Verifikasi (oleh verifikator)
	_, err = db.Exec(`
		INSERT INTO verifikasi (laporan_id, verifikator_id, komentar, status_verified)
		VALUES ($1, $2, 'Sudah dicek dan sesuai requirement', 'approved')
	`, laporanID, verifikatorID)
	if err != nil {
		log.Fatal("Gagal seed verifikasi:", err)
	}
	log.Println("✅ Verifikasi seeded (verifikator: verifikator)")
}

func mustGetUserID(db *sql.DB, username string) string {
	var id string
	err := db.QueryRow(`SELECT id FROM users WHERE username = $1`, username).Scan(&id)
	if err != nil {
		log.Fatalf("User '%s' tidak ditemukan. Pastikan SeedUsers sudah dijalankan: %v", username, err)
	}
	return id
}

func mustGetFirstID(db *sql.DB, table string) string {
	var id string
	err := db.QueryRow(`SELECT id FROM ` + table + ` LIMIT 1`).Scan(&id)
	if err != nil {
		log.Fatalf("Tabel '%s' kosong. Pastikan seeder terkait sudah dijalankan: %v", table, err)
	}
	return id
}
