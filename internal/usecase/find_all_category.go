package usecase

import (
	"context"
	"neuro-most/category-service/internal/entities"
)

type (
	FindAllCategory interface {
		Execute(ctx context.Context, input FindAllCategoryInput) ([]FindAllCategoryOutput, int64, error)
	}

	FindAllCategoryInput struct {
		Page     int64
		PageSize int64
	}

	FindAllCategoryOutput struct {
		Id   int64
		Name string
	}

	FindAllCategoryPresenter interface {
		Output(Categorys []entities.Category) []FindAllCategoryOutput
	}

	findAllCategoryInteractor struct {
		repo      entities.CategoryRepo
		presenter FindAllCategoryPresenter
	}
)

func NewFindAllCategoryInteractor(repo entities.CategoryRepo, presenter FindAllCategoryPresenter) FindAllCategory {
	return &findAllCategoryInteractor{repo: repo, presenter: presenter}
}

func (uc findAllCategoryInteractor) Execute(ctx context.Context, input FindAllCategoryInput) ([]FindAllCategoryOutput, int64, error) {
	categories, total, err := uc.repo.Fetch(ctx, input.Page, input.PageSize)
	if err != nil {
		return nil, 0, err
	}
	return uc.presenter.Output(categories), total, nil
}
