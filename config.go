package main

import (
	"flag"
	//"github.com/nbd-wtf/go-nostr"
	//"github.com/nbd-wtf/go-nostr/nip19"
	"fiatjaf.com/nostr"
	"fiatjaf.com/nostr/nip19"
	"log"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	Port   int
	NSec   string
	NPub   string
	Relays []string
}

func check(e error) {
	if e != nil {
		log.Fatal("Fatal: %s", e)
	}
}

func LoadConfig() *Config {
	var nsecfile string
	var config Config
	var portarg int
	var nsecarg string
	var npubarg string
	var relayargs string

	flag.IntVar(&portarg, "port", 49160, "HTTP listen port. Defaults to 49160/TCP. Can be set through WEBHOOK_PORT env")
	flag.StringVar(&nsecarg, "nsec", "", "Path to file containing bech32-encoded nsec. Can be set through NOSTR_NSEC_FILE env or generated on first run")
	flag.StringVar(&npubarg, "npub", "", "Bech32-encoded npub of the account who will receive notifications. Can be set through NOSTR_NPUB env. You can also set a by alert target by setting the \"npub\" label on the alert properties.")
	flag.StringVar(&relayargs, "relays", "", "Comma separated list of relays you want to send your notifications to. Can be set through NOSTR_RELAYS env.")
	flag.Parse()

	port, ok := os.LookupEnv("WEBHOOK_PORT")
	if ok {
		port, err := strconv.Atoi(port)
		check(err)
		config.Port = port
	} else {
		config.Port = portarg
	}

	nsecvar, ok := os.LookupEnv("NOSTR_NSEC_FILE")
	if ok {
		nsecfile = nsecvar
	}
	if nsecarg != "" {
		nsecfile = nsecarg
	}

	if nsecfile != "" {
		data, err := os.ReadFile(nsecfile)
		check(err)
		nsec := strings.TrimSpace(string(data))
		prefix, decoded, err := nip19.Decode(nsec)
		sk := decoded.(nostr.SecretKey)
		check(err)
		if prefix != "nsec" {
			log.Fatalf("not an nsec, got %s", prefix)
		}
		
		pk := nostr.GetPublicKey(sk)
		npub := nip19.EncodeNpub(pk)
		log.Printf("Loading private key from file %s. Will send alerts from %s", nsecfile, npub)
		config.NSec = nsec
	} else {
		sk := nostr.Generate()
		pk := nostr.GetPublicKey(sk)
		nsec := nip19.EncodeNsec(sk)
		npub := nip19.EncodeNpub(pk)
		log.Printf("Using random private key. Will send alerts from %s", npub)
		config.NSec = nsec
	}

	npub, ok := os.LookupEnv("NOSTR_NPUB")
	if ok {
		config.NPub = npub
	} else {
		config.NPub = npubarg
	}

	relays, ok := os.LookupEnv("NOSTR_RELAYS")
	if ok {
		config.Relays = strings.Split(relays, ",")
	}
	return &config
}
