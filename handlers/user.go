package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/a-h/templ"
	"github.com/asayuki/gopherreads/config"
	"github.com/asayuki/gopherreads/models"
	"github.com/asayuki/gopherreads/stores"
	"github.com/asayuki/gopherreads/templates/components"
	"github.com/asayuki/gopherreads/templates/layout"
	"github.com/asayuki/gopherreads/templates/pages"
)

type UserHandler struct {
	user    *stores.UserStore
	library *stores.LibraryStore
}

func InitUserHandler(user *stores.UserStore, library *stores.LibraryStore) *UserHandler {
	return &UserHandler{
		user,
		library,
	}
}

func (h *UserHandler) AuthView(w http.ResponseWriter, r *http.Request) {
	var t templ.Component
	if r.Header.Get("HX-Request") == "true" {
		t = pages.AuthView()
	} else {
		t = layout.Base(pages.AuthView(), false)
	}

	render(w, r, t, http.StatusOK)
}

func (h *UserHandler) AuthUser(w http.ResponseWriter, r *http.Request) {
	var ua models.UserAuth
	if err := validateFormPayload(r, &ua); err != nil {
		render(w, r, components.Error("Validation error"), http.StatusBadRequest)
		return
	}

	u, err := h.user.GetUserByField("email", ua.Email)
	if err != nil {
		log.Println(err)
		render(w, r, components.Error("Bad credentials"), http.StatusBadRequest)
		return
	}

	if !comparePassword(u.Password, []byte(ua.Password)) {
		render(w, r, components.Error("Bad credentials"), http.StatusBadRequest)
		return
	}

	token, err := createJWT(map[string]interface{}{
		"sub": u.ID,
	}, config.Envs.SessionExp)
	if err != nil {
		log.Println(err)
		render(w, r, components.Error("Bad credentials"), http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "auth_token",
		Value:    token,
		Path:     "/",
		Expires:  time.Now().Add(time.Second * time.Duration(config.Envs.SessionExp)),
		HttpOnly: true,
	})

	w.Header().Set("hx-redirect", "/")
	w.WriteHeader(http.StatusOK)
}
