package service

import (
	"context"

	"github.com/andreoav/mock-helper-api/pkg/domain"
	"github.com/pkg/errors"
)

type ProjectService struct {
	projectRepo domain.ProjectRepository
}

func NewProjectService(projectRepo domain.ProjectRepository) domain.ProjectService {
	return &ProjectService{
		projectRepo: projectRepo,
	}
}

func (ps ProjectService) FetchByBasePath(ctx context.Context, basePath string) (domain.Project, error) {
	project, err := ps.projectRepo.GetByBasePath(ctx, basePath)

	if err != nil {
		return domain.Project{}, errors.Wrap(err, "FetchByBasePath failed")
	}

	return project, nil
}
