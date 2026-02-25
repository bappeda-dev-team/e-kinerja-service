package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	var err error

	dsn := "root:@tcp(127.0.0.1:3306)/db_internal?parseTime=true"

	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Gagal konek DB:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("DB tidak bisa diakses:", err)
	}

	fmt.Println("✅ Database connected!")
}
