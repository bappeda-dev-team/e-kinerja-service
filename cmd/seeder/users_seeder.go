package main

import (
	"database/sql"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func SeedUsers(db *sql.DB) {
	// Ambil semua role berdasarkan nama
	roles := map[string]string{}
	rows, err := db.Query(`SELECT id, name FROM roles WHERE name IN ('super_admin', 'admin', 'programmer', 'verifikator')`)
	if err != nil {
		log.Fatal("Gagal query roles:", err)
	}
	defer rows.Close()
	for rows.Next() {
		var id, name string
		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal("Gagal scan role:", err)
		}
		roles[name] = id
	}

	if len(roles) == 0 {
		log.Fatal("Tidak ada role ditemukan. Jalankan SeedRoles terlebih dahulu.")
	}

	// Daftar user per role
	users := []struct {
		RoleName string
		Username string
		FullName string
		Password string
	}{
		{"super_admin", "pak_yoga", "Pak Yoga", "pakyoga123"},
		{"super_admin", "mas_ndaru", "Mas Ndaru", "masndaru123"},
		{"admin", "mas_ilham", "Mas Ilham", "masilham123"},
		{"verifikator", "mas_ryan", "Mas Ryan", "masryan123"},
		{"programmer", "agnar", "Agnar", "agnar123"},
		{"programmer", "myko", "Myko", "myko123"},
		{"programmer", "affandi", "Affandi", "affandi123"},
		{"programmer", "zulfikar", "Zulfikar", "zulfikar123"},
		{"programmer", "maura", "Maura", "maura123"},
		{"programmer", "daniel", "Daniel", "daniel123"},
	}

	for _, u := range users {
		roleID, ok := roles[u.RoleName]
		if !ok {
			log.Printf("⚠️  Role '%s' tidak ditemukan, skip user '%s'", u.RoleName, u.Username)
			continue
		}

		hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Fatalf("Gagal hash password untuk user '%s': %v", u.Username, err)
		}

		_, err = db.Exec(`
			INSERT INTO users (role_id, username, full_name, password, is_active)
			VALUES ($1, $2, $3, $4, true)
			ON CONFLICT (username) DO NOTHING
		`, roleID, u.Username, u.FullName, string(hashed))
		if err != nil {
			log.Fatalf("Gagal seed user '%s': %v", u.Username, err)
		}

		log.Printf("✅ User '%s' (role: %s) seeded — password: %s", u.Username, u.RoleName, u.Password)
	}
}
