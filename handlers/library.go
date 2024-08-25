package handlers

import (
	"net/http"

	"github.com/asayuki/gopherreads/models"
	"github.com/asayuki/gopherreads/stores"
	"github.com/asayuki/gopherreads/templates/pages"
	"github.com/asayuki/gopherreads/templates/layout"
)

type LibraryHandler struct {
	library *stores.LibraryStore
	user    *stores.UserStore
	cache   map[string]*models.Cache
}

func InitLibraryHandler(library *stores.LibraryStore, user *stores.UserStore, cache map[string]*models.Cache) *LibraryHandler {
	return &LibraryHandler{
		library,
		user,
		cache,
	}
}

func (h *LibraryHandler) BaseView(w http.ResponseWriter, r *http.Request) {
	render(w, r, layout.Base(), 200)
}

func (h *LibraryHandler) CurrentlyReadingView(w http.ResponseWriter, r *http.Request) {
	render(w, r, pages.CurrentlyReadingView(), 200)
}

func (h *LibraryHandler) LibraryView(w http.ResponseWriter, r *http.Request) {

}

func (h *LibraryHandler) OpenBook() {

}
