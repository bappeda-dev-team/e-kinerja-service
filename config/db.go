package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	var err error

	dsn := "host=127.0.0.1 port=5432 user=postgres password=psg2026 dbname=db_internal sslmode=disable"

	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Gagal konek DB:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("DB tidak bisa diakses:", err)
	}

	fmt.Println("✅ Database connected!")
}
