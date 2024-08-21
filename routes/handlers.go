package routes

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/a-h/templ"
	"github.com/asayuki/gopherreads/config"
	"github.com/asayuki/gopherreads/models"
	"github.com/asayuki/gopherreads/stores"
	"github.com/asayuki/gopherreads/templates"
	"github.com/asayuki/gopherreads/templates/features"
)

type Handler struct {
	library     *stores.LibraryStore
	directories *models.Directory
	cache       map[string]*models.Cache
}

func InitHandler(library *stores.LibraryStore, directories *models.Directory, cache map[string]*models.Cache) *Handler {
	return &Handler{library, directories, cache}
}

func (h *Handler) pageView(w http.ResponseWriter, r *http.Request) {
	var t templ.Component

	page := r.URL.Path
	menu := h.directories.GetDirectories(page)

	books, err := h.library.GetBooksByPath(fixURLSlash(page))
	if err != nil {
		fmt.Println(err)
	}

	if r.Header.Get("HX-Request") == "true" {
		t = features.PageView(menu, books)
	} else {
		t = features.Base(features.PageView(menu, books))
	}

	templates.Render(w, r, t, 200)
}

func (h *Handler) openBook(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Query().Get("path")
	fmt.Println(path)
	book, err := h.library.GetBookByPath(path)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(book)
}

func fixURLSlash(url string) string {
	if !strings.HasSuffix(url, "/") {
		url += "/"
	}

	if strings.HasSuffix(config.Envs.LibraryPath, "/") {
		url = strings.TrimPrefix(url, "/")
	}

	return url
}
