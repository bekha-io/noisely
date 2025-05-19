package repository

import (
	"context"
	"noisely/internal/domain/entities"
)

type SchemaRepository interface {
	Create(ctx context.Context, schema *entities.Schema) error
	GetByID(ctx context.Context, id string) (*entities.Schema, error)
	GetByNameAndVersion(ctx context.Context, name string, version string) (*entities.Schema, error)
	GetLatest(ctx context.Context, name string) (*entities.Schema, error)
	List(ctx context.Context, filter map[string]interface{}) ([]*entities.Schema, error)
}
