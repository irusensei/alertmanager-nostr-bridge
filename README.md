# alertmanager-nostr-bridge

Setup a webhook to forward Grafana alerts to Nostr

```
Usage of ./alertmanager-nostr-bridge:
  -npub string
    	Bech32-encoded npub of the account who will receive notifications. Can be set through NOSTR_NPUB env. You can also set a by alert target by setting the "npub" label on the alert properties.
  -nsec string
    	Path to file containing bech32-encoded nsec. Can be set through NOSTR_NSEC_FILE env or generated on first run
  -port int
    	HTTP listen port. Defaults to 49160/TCP. Can be set through WEBHOOK_PORT env (default 49160)
  -relays string
    	Comma separated list of relays you want to send your notifications to. Can be set through NOSTR_RELAYS env.
```
