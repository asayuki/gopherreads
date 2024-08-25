package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/asayuki/gopherreads/config"
	"github.com/asayuki/gopherreads/database"
	"github.com/asayuki/gopherreads/library"
	"github.com/asayuki/gopherreads/middleware"
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
	userStore := stores.NewUserStore(db)

	directories := models.Directory{
		Path: "/",
	}

	library.Categorize(*libraryStore, &directories)

	router := http.NewServeMux()

	// Serve static files
	fs := http.FileServer(http.Dir("./static"))
	router.Handle("/static/", http.StripPrefix("/static/", fs))
	router.Handle("/favicon.ico", http.FileServer(http.Dir("./static")))

	// Register router
	routes.RegisterRoutes(router, *libraryStore, *userStore, cache)

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

	middlewares := middleware.CreateStack(
		middleware.Logging,
	)

	server := http.Server{
		Addr:    config.Envs.HTTPAddr,
		Handler: middlewares(router),
	}

	fmt.Printf("Server started on: http://%v", config.Envs.HTTPAddr)
	server.ListenAndServe()
}
