package service

import (
	"context"

	"github.com/andreoav/mock-helper-api/pkg/domain"
	"github.com/pkg/errors"
)

type MockService struct {
	projectRepo domain.ProjectRepository
}

func NewMockService(projectRepo domain.ProjectRepository) domain.MockService {
	return &MockService{
		projectRepo: projectRepo,
	}
}

func (ms MockService) FetchScenarioByRequest(ctx context.Context, req domain.MockRequest) (domain.Scenario, error) {
	var err error
	var project domain.Project

	if project, err = ms.projectRepo.GetByID(ctx, req.Project); err != nil {
		return domain.Scenario{}, errors.Wrap(err, "FetchScenarioByRequest failed")
	}

	for _, endpoint := range project.Endpoints {
		if endpoint.Method == req.Method && endpoint.Path == req.Path {
			for _, scenario := range endpoint.Scenarios {
				if scenario.Active {
					return scenario, nil
				}
			}
		}
	}

	return domain.Scenario{}, nil
}
