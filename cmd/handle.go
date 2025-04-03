package main

import "net/http"

// Test function for the app
func (app *application) test(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}
