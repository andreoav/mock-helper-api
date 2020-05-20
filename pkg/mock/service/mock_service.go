package service

import (
	"context"

	"github.com/andreoav/mock-helper-api/pkg/domain"
	"github.com/pkg/errors"

	pathToRegexp "github.com/soongo/path-to-regexp"
)

// MockService struct
type MockService struct {
	projectRepo domain.ProjectRepository
}

// NewMockService create a mock value
func NewMockService(projectRepo domain.ProjectRepository) domain.MockService {
	return &MockService{
		projectRepo: projectRepo,
	}
}

// FetchScenarioByRequest servjce
func (ms MockService) FetchScenarioByRequest(ctx context.Context, req domain.MockRequest) (domain.Scenario, error) {
	project, err := ms.projectRepo.GetByID(ctx, req.Project) // TODO: add method filter

	if err != nil {
		return domain.Scenario{}, errors.Wrap(err, "FetchScenarioByRequest failed")
	}

	for _, endpoint := range project.Endpoints {
		if ms.matchEndpointPath(endpoint.Path, req.Path) {
			for _, scenario := range endpoint.Scenarios {
				if scenario.Active {
					return scenario, nil
				}
			}
		}
	}

	return domain.Scenario{}, nil
}

func (ms MockService) matchEndpointPath(targetPath, sourcePath string) bool {
	var tokens []pathToRegexp.Token
	regexp, err := pathToRegexp.PathToRegexp(targetPath, &tokens, nil)

	if err != nil {
		return false
	}

	match, _ := regexp.MatchString(sourcePath)
	return match
}
