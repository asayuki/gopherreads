package routes

import (
	"net/http"

	"github.com/asayuki/gopherreads/models"
	"github.com/asayuki/gopherreads/stores"
	"github.com/asayuki/gopherreads/handlers"
)

func RegisterRoutes(
	router *http.ServeMux,
	library stores.LibraryStore,
	user stores.UserStore,
	cache map[string]*models.Cache,
) {
	// handler := InitHandler(&library, directories, cache)
	handler := handlers.InitLibraryHandler(&library, &user, cache)

	//router.HandleFunc("GET /", handler.pageView)
	//router.HandleFunc("GET /{page}", handler.pageView)
	//router.Handle("/", http.HandlerFunc(handler.dashboardView))
	//router.Handle("/open", http.HandlerFunc(handler.openBook))
	router.Handle("/", http.HandlerFunc(handler.BaseView))
}
