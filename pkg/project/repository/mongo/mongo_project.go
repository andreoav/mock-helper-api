package mongo

import (
	"context"

	"github.com/pkg/errors"

	"github.com/andreoav/mock-helper-api/pkg/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

// GetByID find a project using the primary key
func (pr ProjectRepository) GetByID(ctx context.Context, id string) (domain.Project, error) {
	var project domain.Project

	objectID, _ := primitive.ObjectIDFromHex(id)

	filter := bson.M{
		"_id":              objectID,
		"endpoints.method": "GET",
	}

	err := pr.DB.Collection("projects").FindOne(ctx, filter).Decode(&project)

	if err != nil || err != mongo.ErrNoDocuments {
		return project, errors.Wrap(err, "GetByID failed")
	}

	return project, nil
}

// GetByBasePath retuns a project by its basePath
func (pr ProjectRepository) GetByBasePath(ctx context.Context, basePath string) (domain.Project, error) {
	var project domain.Project

	err := pr.DB.Collection("projects").FindOne(ctx, bson.M{"basePath": basePath}).Decode(&project)

	if err != nil || err != mongo.ErrNoDocuments {
		return project, errors.Wrap(err, "getByBasePath failed")
	}

	return project, nil
}
