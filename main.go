package main

import (
	httpserver "app-server/http-server"
	"log"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(httpserver.PlayerServer)
	log.Fatal(http.ListenAndServe(":5000", handler))
}
