package main

import (
	"log"
	"net/http"
	"time"

	"github.com/asayuki/gopherreads/config"
	"github.com/asayuki/gopherreads/database"
	"github.com/asayuki/gopherreads/library"
	"github.com/asayuki/gopherreads/models"
	"github.com/asayuki/gopherreads/routes"
	"github.com/asayuki/gopherreads/stores"
)

var cache = make(map[string]*models.Cache)

func main() {
	db, err := database.SQLiteStorage(config.Envs.DBFile)
	if err != nil {
		log.Fatal(err)
	}

	libraryStore := stores.NewLibraryStore(db)

	directories := models.Directory{
		Path: "/",
	}

	library.Categorize(*libraryStore, &directories)

	router := http.NewServeMux()

	routes.RegisterRoutes(router, *libraryStore, &directories)

	go func() {
		t := time.NewTicker(1 * time.Minute)
		for range t.C {
			for key, c := range cache {
				if c.IsExpired() {
					delete(cache, key)
				}
			}
		}
	}()

	server := http.Server{
		Addr:    config.Envs.HTTPAddr,
		Handler: router,
	}

	server.ListenAndServe()
}
