package main

import (
	"net/http"

	"github.com/bmizerany/pat"
)

func (app *application) route() http.Handler {
	// A servermux is like a map to for each url to match to each handler
	// Pat allows for a more complex mux giving the ability of stating if the action is get or post
	mux := pat.New()

	// Basic test to see if all flows correctly
	mux.Get("/test", http.HandlerFunc(app.test))

	// Set the root to be home
	mux.Get("/", http.HandlerFunc(app.home))

	// Includes all the static files that will be used and makes them accsesable(CSS, JS, Images)
	fileServe := http.FileServer(http.Dir("./userEnd/static/"))

	// Strips the /static from the request path and maps the file serve to /static/
	mux.Get("/static/", http.StripPrefix("/static", fileServe))

	return mux
}
