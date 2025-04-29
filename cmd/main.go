package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// DB varibles
var (
	db_username string
	db_password string
	db_ip       string
	db_name     string
)

// Struct to pass around different information
// Also helps to get diffrent methods to diffrent files
type application struct {
	infoLog *log.Logger
	errLog  *log.Logger
	DB      *TableGet
}

func main() {
	// Custom error logging to help make more readable
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errLog := log.New(os.Stderr, "ERR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Load env varibles for DB
	err := godotenv.Load("secret/data.env")
	if err != nil {
		errLog.Fatal(err)
	}

	// The DB varibles assigned to the env varibles
	db_username = os.Getenv("DB_USERNAME")
	db_password = os.Getenv("DB_PASSWORD")
	db_ip = os.Getenv("DB_IP")
	db_name = os.Getenv("DB_NAME")

	// Flag to change the port used when running the app
	addr := flag.String("addr", ":8080", "HTTP Network address")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", db_username, db_password, db_ip, db_name)
	flag.Parse()

	// Starts the database connection to the table called todoApp
	fmt.Println("Starting DB connection")
	db, err := openDB(dsn)
	if err != nil {
		errLog.Fatal(err)
	}
	defer db.Close()

	// DB connection pool
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxIdleTime(5 * time.Minute)

	// Creating a app struct
	app := application{
		infoLog: infoLog,
		errLog:  errLog,
		DB:      &TableGet{DB: db},
	}

	// A custom server to listen and serve, Allows for better logging and to choose address
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errLog,
		Handler:  app.route(),
	}

	// Start server
	fmt.Printf("Starting on port %s", *addr)
	err = srv.ListenAndServe()
	log.Fatal(err)
}

// Function to create database connection
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
