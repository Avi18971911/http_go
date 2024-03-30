package main

import (
	poker "hello"
	"log"
	"net/http"
)

func main() {
	store := poker.NewInMemoryPlayerStore()
	server := poker.NewPlayerServer(store)
	router := poker.CreateRouter(server)
	log.Fatal(http.ListenAndServe(":5000", router))
}
