package Handlers

import (
	//"fmt"
	"html/template"
	"net/http"
)

var templates *template.Template
func init() {
	// Load templates
	templates = template.Must(
		template.ParseFiles(
		"templates/index.html",
		"templates/404.html",
		"templates/400.html",
		"templates/405.html",
		"templates/500.html",
		"templates/details.html",
	))
}


func HomePageHandler(w http.ResponseWriter, r *http.Request, artists []FullData) {

if r.URL.Path == "/" {

err := templates.ExecuteTemplate(w, "index.html", artists)
if err != nil {
	InternalServerErrorHandler(w,r)
}
}else{
	NotFoundHandler(w,r)
}
	// fmt.Printf("Data: %+v\n", ArtistsFull)
}
