package main

import (
	"net/http"
)

// Test function for the app
func (app *application) test(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}

// A basic handling for the for the home page
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "userEnd/html/home.html")
}
