package main

import (
	"net/http"
)

// Simple function for hadnling api request
func (app *application) getTables(w http.ResponseWriter, r *http.Request) {
	app.tableHelper(w, r)
}

func (app *application) commodityGet(w http.ResponseWriter, r *http.Request) {
  app.commodityHelper(w, r)
}

func (app *application) resourceGet(w http.ResponseWriter, r *http.Request) {
  app.resourceHelper(w, r)
}
