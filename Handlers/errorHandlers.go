package Handlers

import (
	"net/http"
)

func BadRequestHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	err := templates.ExecuteTemplate(w, "400.html", nil)
	if err != nil {
		InternalServerErrorHandler(w, r)
		return
	}
}
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	err := templates.ExecuteTemplate(w, "404.html", nil)
	if err != nil {
		InternalServerErrorHandler(w, r)
		return
	}
}
func MethodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	err := templates.ExecuteTemplate(w, "405.html", nil)
	if err != nil {
		InternalServerErrorHandler(w, r)
	}
}
func InternalServerErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	err := templates.ExecuteTemplate(w, "500.html", nil)
	if err != nil {
	http.Error(w, "Failed to render template", http.StatusInternalServerError)
}
}


