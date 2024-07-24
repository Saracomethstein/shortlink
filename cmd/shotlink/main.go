package main

import (
	"log"
	"net/http"
	"shotlink/internal/app/handlers"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./website/static")))
	http.HandleFunc("/api", handlers.HandlerHi)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
