package handlers

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/asayuki/gopherreads/models"
	"github.com/asayuki/gopherreads/stores"
	"github.com/asayuki/gopherreads/templates/components"
	"github.com/asayuki/gopherreads/templates/layout"
	"github.com/asayuki/gopherreads/templates/pages"
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
	var t templ.Component
	if r.Header.Get("HX-Request") == "true" {
		t = pages.CurrentlyReadingView()
	} else {
		t = layout.Base(pages.CurrentlyReadingView())
	}

	render(w, r, t, http.StatusOK)
}

func (h *LibraryHandler) CurrentlyReadingView(w http.ResponseWriter, r *http.Request) {
	var t templ.Component
	if r.Header.Get("HX-Request") == "true" {
		t = pages.CurrentlyReadingView()
	} else {
		t = layout.Base(pages.CurrentlyReadingView())
	}

	render(w, r, t, http.StatusOK)
}

func (h *LibraryHandler) LibraryView(w http.ResponseWriter, r *http.Request) {
	var t templ.Component
	if r.Header.Get("HX-Request") == "true" {
		t = pages.LibraryView()
	} else {
		t = layout.Base(pages.LibraryView())
	}

	render(w, r, t, http.StatusOK)
}

func (h *LibraryHandler) OpenBook() {

}

func (h *LibraryHandler) BookView(w http.ResponseWriter, r *http.Request) {
	var t templ.Component
	if r.Header.Get("HX-Request") == "true" {
		t = components.BookView()
	} else {
		t = layout.Base(components.BookView())
	}

	render(w, r, t, http.StatusOK)
}
