package mongo

import (
	"context"
	"noisely/internal/domain/entities"
	"noisely/internal/repository"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type datasetRepository struct {
	collection *mongo.Collection
}

func NewDatasetRepository(db *mongo.Database) repository.DatasetRepository {
	return &datasetRepository{
		collection: db.Collection("datasets"),
	}
}

func (r *datasetRepository) Create(ctx context.Context, dataset *entities.Dataset) error {
	_, err := r.collection.InsertOne(ctx, dataset)
	return err
}

func (r *datasetRepository) GetByID(ctx context.Context, id string) (*entities.Dataset, error) {
	var dataset entities.Dataset
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&dataset)
	return &dataset, err
}

func (r *datasetRepository) List(ctx context.Context, filter map[string]interface{}) ([]*entities.Dataset, error) {
	cur, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var datasets []*entities.Dataset
	if err := cur.All(ctx, &datasets); err != nil {
		return nil, err
	}
	return datasets, nil
}
