package usecase

import (
	"context"
	"neuro-most/category-service/internal/entities"
)

type (
	DeleteCategoryUseCase interface {
		Execute(ctx context.Context, input DeleteCategoryInput) error
	}

	DeleteCategoryInput struct {
		Id int64
	}

	deleteCategoryInteractor struct {
		repo entities.CategoryRepo
	}
)

func NewDeleteCategoryInteractor(repo entities.CategoryRepo) DeleteCategoryUseCase {
	return &deleteCategoryInteractor{repo: repo}
}

func (uc deleteCategoryInteractor) Execute(ctx context.Context, input DeleteCategoryInput) error {
	category, err := uc.repo.GetByID(ctx, input.Id)
	if err != nil {
		return err
	}
	if err := uc.repo.Delete(ctx, category); err != nil {
		return err
	}
	return nil
}
