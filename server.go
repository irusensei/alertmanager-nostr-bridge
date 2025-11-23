package main

import (
	"io"
	"net/http"
)


func (cfg *Config) alertReceiver(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(io.LimitReader(r.Body, 1<<20))
	if err != nil {
		http.Error(w, "Payload exceeded maximum size. Are you really sending alerts?", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if len(body) == 0 {
		http.Error(w, "Empty payload", http.StatusBadRequest)
		return
	}

	payload := ParseAlert(body)
	publish := sendAlerts(cfg, payload)
	check(publish)

}
