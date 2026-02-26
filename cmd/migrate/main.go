package main

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Gunakan: go run cmd/migrate/main.go [up|down]")
	}

	command := os.Args[1]

	// sesuaikan database kamu
	dsn := "postgres://postgres:psg2026@127.0.0.1:5432/db_internal?sslmode=disable"

	m, err := migrate.New(
		"file://migrations",
		dsn,
	)
	if err != nil {
		log.Fatal(err)
	}

	switch command {
	case "up":
		err = m.Up()
		if err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
		fmt.Println("✅ Migration UP sukses")

	case "down":
		err = m.Down()
		if err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
		fmt.Println("✅ Migration DOWN sukses")

	default:
		log.Fatal("Command harus: up atau down")
	}
}
