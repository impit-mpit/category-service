package usecase

import (
	"context"
	"neuro-most/category-service/internal/entities"
)

type (
	UpdateCategoryUseCase interface {
		Execute(ctx context.Context, input UpdateCategoryInput) error
	}

	UpdateCategoryInput struct {
		Id   int64
		Name *string
	}

	updateCategoryInteractor struct {
		repo entities.CategoryRepo
	}
)

func NewUpdateCategoryInteractor(repo entities.CategoryRepo) UpdateCategoryUseCase {
	return &updateCategoryInteractor{repo: repo}
}

func (uc updateCategoryInteractor) Execute(ctx context.Context, input UpdateCategoryInput) error {
	category, err := uc.repo.GetByID(ctx, input.Id)
	if err != nil {
		return err
	}

	if input.Name != nil {
		category.SetName(*input.Name)
	}

	if err := uc.repo.Update(ctx, category); err != nil {
		return err
	}

	return nil
}
