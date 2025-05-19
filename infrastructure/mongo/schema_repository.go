package mongo

import (
	"context"
	"noisely/internal/domain/entities"
	"noisely/internal/repository"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type schemaRepository struct {
	collection *mongo.Collection
}

func NewSchemaRepository(db *mongo.Database) repository.SchemaRepository {
	return &schemaRepository{
		collection: db.Collection("schemas"),
	}
}

func (r *schemaRepository) Create(ctx context.Context, schema *entities.Schema) error {
	_, err := r.collection.InsertOne(ctx, schema)
	return err
}

func (r *schemaRepository) GetByID(ctx context.Context, id string) (*entities.Schema, error) {
	var schema entities.Schema
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&schema)
	return &schema, err
}

func (r *schemaRepository) GetByNameAndVersion(ctx context.Context, name string, version string) (*entities.Schema, error) {
	var schema entities.Schema
	err := r.collection.FindOne(ctx, bson.M{"name": name, "version": version}).Decode(&schema)
	return &schema, err
}

func (r *schemaRepository) GetLatest(ctx context.Context, name string) (*entities.Schema, error) {
	var schema entities.Schema
	opt := options.FindOne().SetSort(bson.D{primitive.E{Key: "version", Value: -1}})
	err := r.collection.FindOne(ctx, bson.M{"name": name}, opt).Decode(&schema)
	return &schema, err
}

func (r *schemaRepository) List(ctx context.Context, filter map[string]interface{}) ([]*entities.Schema, error) {
	cur, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var schemas []*entities.Schema
	if err := cur.All(ctx, &schemas); err != nil {
		return nil, err
	}
	return schemas, nil
}
