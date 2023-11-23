
package Handlers

import (
	"io/ioutil"
	"log"
	"encoding/json"
	"net/http"
)

func fetchData(url string, data interface{}) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	if response.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		// Unmarshal JSON data into the provided data interface
		err = json.Unmarshal(body, &data)
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Printf("Data: %+v\n", data)
	}
}
