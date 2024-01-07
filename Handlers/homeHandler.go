package Handlers

import (
	//"fmt"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

var templates *template.Template
func init() {
// Load templates
var err error
templates, err = template.ParseFiles(
	"templates/inde.html",
	"templates/404.html",
	"templates/400.html",
	"templates/405.html",
	"templates/500.html",
	"templates/details.html",
)
if err != nil {
	// Print the detailed error message
	fmt.Println("Failed to parse templates:", err)
	os.Exit(1)
}


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
