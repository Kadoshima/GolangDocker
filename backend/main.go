package main

import (
	"backend/infrastructure/router"
	"log"
	"net/http"
)

func main() {
	r := router.SetupRouter()

	log.Println("Starting server on :8000")
	if err := http.ListenAndServe(":8000", r); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
