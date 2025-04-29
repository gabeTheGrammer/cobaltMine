package main

import (
	"net/http"
)

// A basic handling for the for the home page
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "userEnd/html/home.html")
}

// A basic handling for the login page
func (app *application) userChoice(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "userEnd/html/userChoice.html")
}
