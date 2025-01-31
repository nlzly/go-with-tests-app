package main

import (
	httpserver "app-server/http-server"
	"log"
	"net/http"
	"os"
)

const dbFileName = "game.db.json"

func main() {
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}

	store := httpserver.NewFileSystemPlayerStore(db)
	server := httpserver.NewServer(store)
	log.Fatal(http.ListenAndServe(":5000", server))
}
