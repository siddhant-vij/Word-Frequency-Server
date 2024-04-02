package main

import (
	"log"
	"net/http"

	"github.com/siddhant-vij/Word-Frequency-Server/internal/httpapi"
)

func main() {
	http.HandleFunc("/frequency", httpapi.HandleFrequencyRequest)
	log.Println("Server is listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Example Usage:
// http://localhost:8080/frequency?word=text&numServerThreads=12
