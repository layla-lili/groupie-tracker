package main 

import (
	"net/http"
	"groupie/Handlers"
	)


func main() {

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", Handlers.HomePageHandler)
	// Handle not found: 404 page 
	http.HandleFunc("/404", Handlers.NotFoundHandler)
	// Handle Bad Request : 400
	http.HandleFunc("/400", Handlers.BadRequestHandler)
	// Handle Bad Request : 405
	http.HandleFunc("/405", Handlers.MethodNotAllowedHandler)
	// Handle Server error: 500
	http.HandleFunc("/500", Handlers.InternalServerErrorHandler)
	http.ListenAndServe(":8080", nil)
}