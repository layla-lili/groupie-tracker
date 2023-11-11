package Handlers

import (
	"net/http"
	"html/template"
)

var templates *template.Template
func init() {
	// Load templates
	templates = template.Must(template.ParseFiles(
		"templates/index.html",
		"templates/404.html",
		"templates/400.html",
		"templates/500.html",
	))
}
func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	url := "https://groupietrackers.herokuapp.com/api/artists"
	artists := Artists{}
	fetchData(url, &artists)
	err := templates.ExecuteTemplate(w, "index.html", artists)
	if err != nil {
		//http.Error(w, "Failed to render template", http.StatusInternalServerError)
		return
	}
}