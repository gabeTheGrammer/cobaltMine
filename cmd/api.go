package main

import (
	"net/http"
)

// Simple function for hadnling api request
func (app *application) getTables(w http.ResponseWriter, r *http.Request) {
	app.tableHelper(w, r)
}
