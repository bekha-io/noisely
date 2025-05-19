package usecase

import (
	"context"
	"noisely/internal/domain/entities"
	"noisely/internal/repository"
)

type DatasetUsecase struct {
	repo repository.DatasetRepository
}

func NewDatasetUsecase(repo repository.DatasetRepository) *DatasetUsecase {
	return &DatasetUsecase{repo: repo}
}

func (u *DatasetUsecase) Create(ctx context.Context, dataset *entities.Dataset) error {
	return u.repo.Create(ctx, dataset)
}

func (u *DatasetUsecase) GetByID(ctx context.Context, id string) (*entities.Dataset, error) {
	return u.repo.GetByID(ctx, id)
}

func (u *DatasetUsecase) List(ctx context.Context, filter map[string]interface{}) ([]*entities.Dataset, error) {
	return u.repo.List(ctx, filter)
}
