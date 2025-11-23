package main

import (
	"strings"
	"encoding/json"
)

func ParseAlert(payload []byte) (string) {
	var data Payload
	p1 := json.Unmarshal(payload, &data)
	check(p1)

	var o strings.Builder

	o.WriteString(data.Title)
	o.WriteString("status: " + data.Status)
	o.WriteString("state: " + data.State)
	o.WriteString(data.Message)
	return o.String()
}
