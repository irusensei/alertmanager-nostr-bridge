package main

import (
	"context"
	"fiatjaf.com/nostr"
	"log"
	"time"
)

func sendAlerts(config *Config, alert string) error {

	pool := nostr.NewPool(nostr.PoolOptions{
		PenaltyBox: true,
	})
	for _, relayurl := range config.Relays {
		if !nostr.IsValidRelayURL(relayurl) {
			log.Fatalf("%s is not a valid relay URL.")
		}
		relay, err := pool.EnsureRelay(relayurl)
		check(err)
		log.Printf("Connected to relay: %s", relay.URL)
	}
	event := &nostr.Event{
		CreatedAt: nostr.Now(),
		Kind:      nostr.KindTextNote,
		Tags:      nostr.Tags{},
		Content:   alert,
	}

	event.Sign(config.SecretKey)
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	results := pool.PublishMany(ctx, config.Relays, *event)

	successCount := 0
	for result := range results {
		if result.Error != nil {
			log.Printf("Failed on %s: %v\n", result.RelayURL, result.Error)
		} else {
			log.Printf("Succeeded on %s\n", result.RelayURL)
			successCount++
		}
	}

	return nil
}
