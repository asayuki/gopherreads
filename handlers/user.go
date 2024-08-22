package handlers

import (
	"net/http"

	"github.com/asayuki/gopherreads/models"
	"github.com/asayuki/gopherreads/stores"
	"github.com/asayuki/gopherreads/templates/components"
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

func (h *UserHandler) AuthUser(w http.ResponseWriter, r *http.Request) {
	var u models.UserAuth
	if err := validateFormPayload(r, &u); err != nil {
		render(w, r, components.Error("Validation error"), http.StatusBadRequest)
	}
}
