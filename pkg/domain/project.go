package domain

import "context"

// Project struct
type Project struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	BasePath  string     `json:"basePath"`
	Endpoints []Endpoint `json:"endpoints"`
}

// ProjectRepository interface
type ProjectRepository interface {
	GetByBasePath(ctx context.Context, basePath string) (Project, error)
}

// ProjectService interface
type ProjectService interface {
	FetchByBasePath(ctx context.Context, basePath string) (Project, error)
}
