package main

import (
	"groupie/Handlers"
	"net/http"
)

const (
	urlArtists   = "https://groupietrackers.herokuapp.com/api/artists"
	urlLocations = "https://groupietrackers.herokuapp.com/api/locations"
	urlDates     = "https://groupietrackers.herokuapp.com/api/dates"
	urlRelation  = "https://groupietrackers.herokuapp.com/api/relation"
)

var ArtistsFull = make([]Handlers.FullData, 0, len(Handlers.Artists{}))

func main() {

	artists := Handlers.Artists{}
	locations := Handlers.Locations{}
	dates := Handlers.Dates{}
	Relation := Handlers.RelationsData{}

	Handlers.FetchData(urlArtists, &artists)
	Handlers.FetchData(urlLocations, &locations)
	Handlers.FetchData(urlDates, &dates)
	Handlers.FetchData(urlRelation, &Relation)

	for i, artist := range artists {

		tmpl := Handlers.FullData{
			ID:      artists[i].ID,
			Image:   artists[i].Image,
			Name:    artists[i].Name,
			Members: make(map[string]string),

			CreationDate:   artists[i].CreationDate,
			FirstAlbum:     artists[i].FirstAlbum,
			Locations:      locations.Index[i].Locations,
			Dates:          dates.Index[i].Dates,
			DatesLocations: Relation.Index[i].DatesLocations,
		}
		for _, member := range artists[i].Members {
			// Set the member name as both the key and value in the map
			tmpl.Members[member] = member
		}

		if tmpl.Image == "https://groupietrackers.herokuapp.com/api/images/mamonasassassinas.jpeg" {
			artist.Image = "static/Images/ops.jpg"
		}

		ArtistsFull = append(ArtistsFull, tmpl)
	}

	staticDir := http.Dir("static")
	fs := http.FileServer(staticDir)
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Register the handlers
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		Handlers.HomePageHandler(w, r, ArtistsFull)
	})

	http.HandleFunc("/details", func(w http.ResponseWriter, r *http.Request) {
		Handlers.DetailspageHandler(w, r, ArtistsFull)
	})

	http.HandleFunc("/404", Handlers.NotFoundHandler)
	http.HandleFunc("/400", Handlers.BadRequestHandler)
	http.HandleFunc("/405", Handlers.MethodNotAllowedHandler)
	http.HandleFunc("/500", Handlers.InternalServerErrorHandler)

	// Start the server
	http.ListenAndServe(":8080", nil)
}
