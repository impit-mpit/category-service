package usecase

import (
	"context"
	"neuro-most/category-service/internal/entities"
)

type (
	CreateCategoryUseCase interface {
		Execute(ctx context.Context, input CreateCategoryInput) error
	}

	CreateCategoryInput struct {
		Name string
	}

	createCategoryInteractor struct {
		repo entities.CategoryRepo
	}
)

func NewCreateCategoryInteractor(repo entities.CategoryRepo) CreateCategoryUseCase {
	return &createCategoryInteractor{repo: repo}
}

func (uc createCategoryInteractor) Execute(ctx context.Context, input CreateCategoryInput) error {
	category := entities.NewCategoryCreate(input.Name)
	if err := uc.repo.Create(ctx, category); err != nil {
		return err
	}
	return nil
}
