package main

import (
	httpserver "app-server/http-server"
	"log"
	"net/http"
)

func main() {
	store := httpserver.NewInMemoryPlayerStore()
	server := httpserver.NewServer(store)
	log.Fatal(http.ListenAndServe(":5000", server))
}
