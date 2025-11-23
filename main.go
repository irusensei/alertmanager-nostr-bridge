package main

import (
	"log"
	"fmt"
	"net/http"
)

func main() {
	cfg := LoadConfig()
	http.HandleFunc("/", cfg.alertReceiver)
	log.Printf("Server starting on http://localhost:%d", cfg.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), nil))
}
