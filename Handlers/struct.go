package Handlers

import (
	"log"
	"io/ioutil"
	"fmt"
	"net/http"
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

func fetch(url string){
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
		fmt.Println(string(body))
	}
}