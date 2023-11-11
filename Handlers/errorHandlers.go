package Handlers

import (
	"net/http"
	"html/template"
)

func BadRequestHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	err := templates.ExecuteTemplate(w, "400.html", nil)
	if err != nil {
		internalServerErrorHandler(w, r)
		return
	}
}
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	err := templates.ExecuteTemplate(w, "404.html", nil)
	if err != nil {
		internalServerErrorHandler(w, r)
		return
	}
}
func internalServerErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	templates.ExecuteTemplate(w, "500.html", nil)
}