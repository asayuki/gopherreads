package routes

import (
	"net/http"

	"github.com/asayuki/gopherreads/models"
	"github.com/asayuki/gopherreads/stores"
)

func RegisterRoutes(
	router *http.ServeMux,
	library stores.LibraryStore,
	directories *models.Directory,
	cache map[string]*models.Cache,
) {
	handler := InitHandler(&library, directories, cache)

	//router.HandleFunc("GET /", handler.pageView)
	//router.HandleFunc("GET /{page}", handler.pageView)
	router.Handle("/", http.HandlerFunc(handler.pageView))
}
