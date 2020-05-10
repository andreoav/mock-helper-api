package domain

// Scenario struct
type Scenario struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Active      bool   `json:"active"`
	Response    string `json:"response"`
}
