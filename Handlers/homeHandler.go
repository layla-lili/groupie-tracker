package Handlers

import (
	//"fmt"
	"html/template"
	"net/http"
)

var templates *template.Template
var ArtistsFull []FullData // Move ArtistsFull to package-level variable

func init() {
	// Load templates
	templates = template.Must(template.ParseFiles(
		"templates/index.html",
		"templates/404.html",
		"templates/400.html",
		"templates/405.html",
		"templates/500.html",
		"templates/details.html",
	))
// Register the /details route handler
http.HandleFunc("/details", func(w http.ResponseWriter, r *http.Request) {
	detailspageHandler(w, r, ArtistsFull)
})


}

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	//https://mholt.github.io/json-to-go/
	urlArtists := "https://groupietrackers.herokuapp.com/api/artists"
	urlLocations := "https://groupietrackers.herokuapp.com/api/locations"
	urlDates := "https://groupietrackers.herokuapp.com/api/dates"
	urlRelation := "https://groupietrackers.herokuapp.com/api/relation"

	artists := Artists{}
	locations := Locations{}
	dates := Dates{}
	Relation := RelationsData{}

	fetchData(urlArtists, &artists)
	fetchData(urlLocations, &locations)
	fetchData(urlDates, &dates)
	fetchData(urlRelation, &Relation)

	ArtistsFull = nil // Clear ArtistsFull before populating it again

	for i := range artists {
		var tmpl FullData
		tmpl.ID = artists[i].ID
		if artists[i].Image == "https://groupietrackers.herokuapp.com/api/images/mamonasassassinas.jpeg" {
			artists[i].Image = "static/Images/ops.jpg"
		}
		tmpl.Image = artists[i].Image
		tmpl.Name = artists[i].Name
		tmpl.Members = artists[i].Members
		tmpl.CreationDate = artists[i].CreationDate
		tmpl.FirstAlbum = artists[i].FirstAlbum
		tmpl.Locations = locations.Index[i].Locations
		tmpl.Dates = dates.Index[i].Dates
		tmpl.DatesLocations = Relation.Index[i].DatesLocations
		ArtistsFull = append(ArtistsFull, tmpl)
	}

err := templates.ExecuteTemplate(w, "index.html", ArtistsFull)
if err != nil {
	http.Error(w, "Failed to render template", http.StatusInternalServerError)
	return
}



	// fmt.Printf("Data: %+v\n", ArtistsFull)
	

}
