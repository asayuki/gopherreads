package routes

import (
	"net/http"

	"github.com/asayuki/gopherreads/handlers"
	"github.com/asayuki/gopherreads/middleware"
	"github.com/asayuki/gopherreads/models"
	"github.com/asayuki/gopherreads/stores"
)

func RegisterRoutes(
	router *http.ServeMux,
	library stores.LibraryStore,
	user stores.UserStore,
	cache map[string]*models.Cache,
) {
	authMiddleware := middleware.Auth(&user)

	// handler := InitHandler(&library, directories, cache)
	libhandler := handlers.InitLibraryHandler(&library, &user, cache)
	userhandler := handlers.InitUserHandler(&user, &library)

	// Library
	router.Handle("/", authMiddleware(http.HandlerFunc(libhandler.BaseView)))
	router.Handle("/library", authMiddleware(http.HandlerFunc(libhandler.LibraryView)))

	// Books
	router.Handle("/book/{id}", authMiddleware(http.HandlerFunc(libhandler.BookView)))

	// Auth stuff
	router.Handle("/auth", http.HandlerFunc(userhandler.AuthView))
	router.HandleFunc("POST /auth", userhandler.AuthUser)
}
