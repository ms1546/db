package main

import (
	"log"
	"net/http"

	"pq/internal/app/config"
	"pq/internal/app/handler"
	"pq/pkg/db"

	_ "github.com/lib/pq"
)

func main() {
	cfg := config.Load()

	dbConn, err := db.InitDB(cfg.Database)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer dbConn.Close()

	h := handler.NewHandler(dbConn)
	http.Handle("/", h.Router())

	log.Printf("Starting server on %s\n", cfg.Server.Address)
	if err := http.ListenAndServe(cfg.Server.Address, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
