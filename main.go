package main

import (
	"log"
	"net/http"
	"test-go/internal/handlers"
)

func main() {
	http.HandleFunc("/notes", handlers.HandleNotes)

	log.Println("test 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("not start server: $s\n", err.Error())
	}
}