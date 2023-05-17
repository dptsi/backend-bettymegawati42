package main

import (
	"flag"
	"log"
	"net/http"

	"demo/internal/models"
)

type application struct {
	readinglist *models.ReadinglistModel
}

func main() {
	addr := flag.String("addr", ":80", "HTTP network address")
	endpoint := flag.String("endpoint", "http://localhost:3000/v1/books", "Endpoint for the readinglist web service")

	app := &application{
		readinglist: &models.ReadinglistModel{Endpoint: *endpoint}, //ditaruh endpoint jadi yg reading list semua mengarah ke satu URL
	}

	srv := &http.Server{
		Addr:    *addr,
		Handler: app.routes(),
	}

	log.Printf("Starting the server on %s", *addr)
	err := srv.ListenAndServe()
	log.Fatal(err)
}
