package templates

import (
	"net/http"

	"github.com/a-h/templ"
)

func Render(w http.ResponseWriter, r *http.Request, template templ.Component, status int) error {
	w.WriteHeader(status)
	return template.Render(r.Context(), w)
}
