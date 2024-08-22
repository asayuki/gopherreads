package handlers

import (
	"errors"
	"net/http"
	"reflect"
	"strings"

	"github.com/a-h/templ"
	"github.com/go-playground/form"
)

var formdecoder = form.NewDecoder()

func validateFormPayload(r *http.Request, dst interface{}) error {
	if err := r.ParseForm(); err != nil {
		return err
	}

	if err := formdecoder.Decode(dst, r.Form); err != nil {
		return err
	}

	structFields := getStructFields(dst)
	for key := range r.Form {
		if _, ok := structFields[key]; !ok {
			return errors.New("unexpected field: " + key)
		}
	}

	return nil
}

func getStructFields(dst interface{}) map[string]struct{} {
	fields := make(map[string]struct{})
	val := reflect.ValueOf(dst).Elem()
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		tag := field.Tag.Get("schema")
		if tag != "" && tag != "-" {
			fields[tag] = struct{}{}
		} else {
			fields[strings.ToLower(field.Name)] = struct{}{}
		}
	}

	return fields
}

func render(w http.ResponseWriter, r *http.Request, template templ.Component, status int) error {
	w.WriteHeader(status)
	return template.Render(r.Context(), w)
}
