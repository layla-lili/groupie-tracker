package main 

import (
	"net/http"
	//"html/template"
	"groupie/Handlers"
	)


func main() {
	http.HandleFunc("/", Handlers.HomePageHandler)
	http.ListenAndServe(":8080", nil)
}