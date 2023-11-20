package Handlers

import (
	"log"
	"io/ioutil"
	//"fmt"
	"net/http"
	"encoding/json"
	)
	type FullData struct {
		ID           int      `json:"id"`
		Image        string   `json:"image"`
		Name         string   `json:"name"`
		Members      []string `json:"members"`
		CreationDate int      `json:"creationDate"`
		FirstAlbum   string   `json:"firstAlbum"`
		Locations []string `json:"locations"`
		Dates     []string   `json:"dates"`
		DatesLocations map[string][]string `json:"datesLocations"`
	}

type Artists []struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	//Locations     interface{}    
	///ConcertDates string   `json:"concertDates"`
	//Relations    string   `json:"relations"`
}

type Locations struct {
	Index []struct {
		ID        int      `json:"id"`
		Locations []string `json:"locations"`
		//Dates     string   `json:"dates"`
	} `json:"index"`
}

type Dates struct {
	Index []struct {
		ID        int      `json:"id"`
		Dates     []string   `json:"dates"`
	} `json:"index"`
}

type RelationsData struct {
	Index []struct {
		DatesLocations map[string][]string `json:"datesLocations"`
	} `json:"index"`
}

func fetchData(url string,data interface{}){
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
        // Unmarshal JSON data into the provided data interface
        err = json.Unmarshal(body, &data)
        if err != nil {
            log.Fatal(err)
        }
		//fmt.Printf("Data: %+v\n", data)

	}
}