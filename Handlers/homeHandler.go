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
	//https://mholt.github.io/json-to-go/
	//urlArtists := "https://groupietrackers.herokuapp.com/api/artists"
	urlLocations:= "https://groupietrackers.herokuapp.com/api/locations"
	//urlDates:=	"https://groupietrackers.herokuapp.com/api/dates"
	//urlRelation:=	"https://groupietrackers.herokuapp.com/api/relation"
	artists := Artists{}
	//fetchData(urlArtists, &artists)
	fetch(urlLocations)

	err := templates.ExecuteTemplate(w, "index.html", artists)
	if err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
		return
	}
}