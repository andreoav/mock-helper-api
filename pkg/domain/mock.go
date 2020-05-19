package domain

import "context"

type MockRequest struct {
	Project string `json:"project"`
	Method  string `json:"method"`
	Path    string `json:"path"`
}

type MockService interface {
	FetchScenarioByRequest(ctx context.Context, req MockRequest) (Scenario, error)
}
