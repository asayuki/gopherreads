package main

import (
	"fmt"
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

	// Some static
	router.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
	})

	// Register other routes
	routes.RegisterRoutes(router, *libraryStore, &directories, cache)

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

	fmt.Println("Top-level directories:")
	topDirs := directories.GetDirectories("/")
	for _, dir := range topDirs {
		fmt.Println(dir.Path, dir.Name)
	}

	fmt.Println("\nSubdirectories under 'Programming':")
	progDirs := directories.GetDirectories("/Programming")
	for _, dir := range progDirs {
		fmt.Println(dir.Path, dir.Name)
	}

	fmt.Println("\nSubdirectories under 'Programming':")
	goDirs := directories.GetDirectories("/Programming/Go")
	for _, dir := range goDirs {
		fmt.Println(dir.Path, dir.Name)
	}

	server := http.Server{
		Addr:    config.Envs.HTTPAddr,
		Handler: router,
	}

	fmt.Printf("Server started on: http://%v", config.Envs.HTTPAddr)
	server.ListenAndServe()
}
