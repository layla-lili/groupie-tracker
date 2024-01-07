package Handlers

import (
	"net/http"
	"strconv"
)



func DetailspageHandler(w http.ResponseWriter, r *http.Request, artists []FullData){
	// Get the selected ID from the form data
	idStr := r.FormValue("id")
	id, err := strconv.Atoi(idStr)
	
	if err != nil {
		// If no artist is found, return a 404 Not Found response
		// http.NotFound(w, r)
		NotFoundHandler(w,r)
	}
	if  id < 1 || id > 52 {
		// Handle the error
		// http.NotFound(w, r)
		BadRequestHandler(w,r)
	}

	// Fetch the artist details using the selected ID
	artist := getArtistDetails(id, artists)
	// Render the details.html template with the artist data
	// tmpl := template.Must(template.ParseFiles("templates/details.html"))
	// tmpl.Execute(w, artist)

error := templates.ExecuteTemplate(w, "details.html", artist)
if error != nil {
	InternalServerErrorHandler(w,r)
}


}

// Function to fetch artist details based on the ID
func getArtistDetails(id int, artists []FullData) FullData {
	// Find the artist with the matching ID
	for _, a := range artists {
		if a.ID == id {
			return a
		}
	}

	// If no artist is found, return an empty FullData struct
	return FullData{}
}
