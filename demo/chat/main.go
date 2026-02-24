package main

import (
	"log"
	"net/http"

	"github.com/troublete/go-lobby/server"
)

func main() {
	http.Handle("/chat/", http.StripPrefix("/chat/", http.FileServer(http.Dir("./internal/public/chat"))))
	http.Handle("/", server.NewLobbyMux())
	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		log.Fatal(err)
	}
}
