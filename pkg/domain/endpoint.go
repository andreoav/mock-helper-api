package domain

// Endpoint struct
type Endpoint struct {
	Path      string     `json:"path"`
	Method    string     `json:"method"`
	Scenarios []Scenario `json:"scenarios"`
}
