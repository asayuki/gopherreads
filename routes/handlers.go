package routes

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/asayuki/gopherreads/models"
	"github.com/asayuki/gopherreads/stores"
	"github.com/asayuki/gopherreads/templates"
	"github.com/asayuki/gopherreads/templates/features"
)

type Handler struct {
	library     *stores.LibraryStore
	directories *models.Directory
}

func InitHandler(library *stores.LibraryStore, directories *models.Directory) *Handler {
	return &Handler{library, directories}
}

func (h *Handler) indexView(w http.ResponseWriter, r *http.Request) {
	var t templ.Component
	if r.Header.Get("HX-Request") == "true" {
		t = features.IndexView(h.directories)
	} else {
		t = features.Base(features.IndexView(h.directories))
	}

	templates.Render(w, r, t, 200)
}
