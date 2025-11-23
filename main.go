package main

import (
	"log"
)

func main() {
	cfg := LoadConfig()
	log.Printf("Port: %s", cfg.Port)
	log.Printf("Nsec: %s", cfg.NSec)
	log.Printf("Npub: %s", cfg.NPub)
	log.Printf("Relays: %s", cfg.Relays)
}
