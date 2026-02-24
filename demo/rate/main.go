package main

import (
	"log"
	"net/http"

	"github.com/troublete/go-lobby/server"
)

func main() {
	http.Handle("/rate/", http.StripPrefix("/rate/", http.FileServer(http.Dir("./internal/public/rate"))))
	http.Handle("/", server.NewLobbyMux())
	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		log.Fatal(err)
	}
}
