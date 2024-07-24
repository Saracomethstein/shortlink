package main

import (
	"log"
	"net/http"
	"shotlink/internal/app/handlers"
)

func main() {
	http.HandleFunc("/", handlers.HandlerHi)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
