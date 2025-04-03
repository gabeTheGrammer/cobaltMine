package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

// Struct to pass around different information
// Also helps to get diffrent methods to diffrent files
type application struct {
	infoLog *log.Logger
	errLog  *log.Logger
}

func main() {
	// Flag to change the port used when running the app
	addr := flag.String("addr", ":8080", "HTTP Network address")

	// Custom error logging to help make more readable
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errLog := log.New(os.Stderr, "ERR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Creating a app struct
	app := application{
		infoLog: infoLog,
		errLog:  errLog,
	}

	// A custom server to listen and serve, Allows for better logging and to choose address
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errLog,
		Handler:  app.route(),
	}

	fmt.Println("Starting on port :8080")
	err := srv.ListenAndServe()
	log.Fatal(err)
}
