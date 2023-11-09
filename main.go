package main 

import (
	"log"
	"io/ioutil"
	"fmt"
	"net/http"
	"html/template"
	"encoding/json"


)


type Artists []struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	//Locations    string   `json:"locations"`
	///ConcertDates string   `json:"concertDates"`
	//Relations    string   `json:"relations"`
}
var templates *template.Template
func init() {
	// Load templates
	templates = template.Must(template.ParseFiles(
		"index.html",
		//"templates/404.html",
		//"templates/400.html",
		//"templates/500.html",
	))
}

func main() {
	http.HandleFunc("/", HomePageHandler)
	http.ListenAndServe(":8080", nil)
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

func fetchData(url string,data *Artists){
	response, err :=http.Get(url)

	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	if response.StatusCode ==http.StatusOK {
		body, err :=ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Println(string(body))
		json.Unmarshal(body,&data)
		fmt.Printf("Data: %+v", data)
	}
	
}