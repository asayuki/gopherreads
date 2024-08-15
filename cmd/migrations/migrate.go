package main

import (
	"log"
	"os"

	"github.com/asayuki/gopherreads/config"
	"github.com/asayuki/gopherreads/database"
	"github.com/golang-migrate/migrate/v4"
	sqlite3Migrate "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	if config.Envs.DBFile == "" {
		log.Fatal("no db specified")
		return
	}

	db, err := database.SQLiteStorage(config.Envs.DBFile)
	if err != nil {
		log.Fatal(err)
	}

	driver, err := sqlite3Migrate.WithInstance(db, &sqlite3Migrate.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://database/migrations", "sqlite3", driver)
	if err != nil {
		log.Fatal(err)
	}

	v, d, _ := m.Version()
	log.Printf("database migrations version: %d, dirty: %v", v, d)

	cmd := os.Args[len(os.Args)-1]
	if cmd == "up" {
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}
	if cmd == "down" {
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatal(err)
		}
	}
}
