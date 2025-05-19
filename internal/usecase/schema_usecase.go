package usecase

import (
	"context"
	"noisely/internal/domain/entities"
	"noisely/internal/repository"
)

type SchemaUsecase struct {
	repo repository.SchemaRepository
}

func NewSchemaUsecase(repo repository.SchemaRepository) *SchemaUsecase {
	return &SchemaUsecase{repo: repo}
}

func (u *SchemaUsecase) Create(ctx context.Context, schema *entities.Schema) error {
	return u.repo.Create(ctx, schema)
}

func (u *SchemaUsecase) GetByID(ctx context.Context, id string) (*entities.Schema, error) {
	return u.repo.GetByID(ctx, id)
}

func (u *SchemaUsecase) GetByNameAndVersion(ctx context.Context, name, version string) (*entities.Schema, error) {
	return u.repo.GetByNameAndVersion(ctx, name, version)
}

func (u *SchemaUsecase) GetLatest(ctx context.Context, name string) (*entities.Schema, error) {
	return u.repo.GetLatest(ctx, name)
}

func (u *SchemaUsecase) List(ctx context.Context, filter map[string]interface{}) ([]*entities.Schema, error) {
	return u.repo.List(ctx, filter)
}
