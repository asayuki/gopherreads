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
) {
	handler := InitHandler(&library, directories)

	router.HandleFunc("GET /", handler.indexView)
}
