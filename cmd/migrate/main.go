package main

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if len(os.Args) < 2 {
		log.Fatal("Gunakan: go run cmd/migrate/main.go [up|down]")
	}

	command := os.Args[1]

	// sesuaikan database kamu
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)

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
