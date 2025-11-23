package main


type Alert struct {
	Status       string         `json:"status"`
	StartsAt     string         `json:"startsAt"`
	EndsAt       string         `json:"endsAt"`
	GeneratorURL string         `json:"generatorURL"`
	Fingerprint  string         `json:"fingerprint"`
	SilenceURL   string         `json:"silenceURL"`
	DashboardURL string         `json:"dashboardURL"`
	PanelURL     string         `json:"panelURL"`
	ValueString  string         `json:"valueString"`
	OrgId        int            `json:"orgId"`
	Labels       map[string]any `json:"labels"`
	Annotations  map[string]any `json:"annotations"`
	values       map[int]any    `json:"values"`
}

type Payload struct {
	Receiver          string         `json:"receiver"`
	Status            string         `json:"firing"`
	ExternalURL       string         `json:"externalURL"`
	Version           string         `json:"version"`
	GroupKey          string         `json:"groupKey"`
	TruncatedAlerts   int            `json:"truncatedAlerts"`
	OrgID             int            `json:"orgId"`
	Title             string         `json:"title"`
	State             string         `json"state"`
	Message           string         `json:"message"`
	Alerts            []Alert        `json:"alerts"`
	GroupLabels       map[string]any `json:"groupLabels"`
	CommonLabels      map[string]any `json:"groupLabels"`
	CommonAnnotations map[string]any `json:"commonAnnotations"`
}
