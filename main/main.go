package main

import (
	"log"
	"net/http"
)

func main() {
	store := NewInMemoryPlayerStore()
	server := NewPlayerServer(store)
	router := createRouter(server)
	log.Fatal(http.ListenAndServe(":5000", router))
}
