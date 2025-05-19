package repository

import (
	"context"
	"noisely/internal/domain/entities"
)

type DatasetRepository interface {
	Create(ctx context.Context, dataset *entities.Dataset) error
	GetByID(ctx context.Context, id string) (*entities.Dataset, error)
	List(ctx context.Context, filter map[string]interface{}) ([]*entities.Dataset, error)
}
