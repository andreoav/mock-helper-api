package mongo

import (
	"context"

	"github.com/pkg/errors"

	"github.com/andreoav/mock-helper-api/pkg/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// ProjectRepository mongo struct
type ProjectRepository struct {
	DB *mongo.Database
}

// NewProjectRepository create a configure mongo repository
func NewProjectRepository(db *mongo.Database) domain.ProjectRepository {
	return &ProjectRepository{
		DB: db,
	}
}

// GetByBasePath retuns a project by its basePath
func (pr ProjectRepository) GetByBasePath(ctx context.Context, basePath string) (domain.Project, error) {
	var project domain.Project

	filter := bson.M{"basePath": basePath}
	err := pr.DB.Collection("projects").FindOne(ctx, filter).Decode(&project)

	if err != nil {
		return domain.Project{}, errors.Wrap(err, "getByBasePath failed")
	}

	return project, nil
}
