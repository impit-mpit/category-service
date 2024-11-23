package usecase

import (
	"context"
	"neuro-most/category-service/internal/entities"
)

type (
	GetByIdCategoryUseCase interface {
		Execute(ctx context.Context, input GetByIdCategoryInput) (GetByIdCategoryOutput, error)
	}

	GetByIdCategoryInput struct {
		Id int64
	}

	GetByIdCategoryOutput struct {
		Id   int64
		Name string
	}

	GetByIdCategoryPresenter interface {
		Output(Category entities.Category) GetByIdCategoryOutput
	}

	getByIdCategoryInteractor struct {
		repo      entities.CategoryRepo
		presenter GetByIdCategoryPresenter
	}
)

func NewGetByIdCategoryInteractor(repo entities.CategoryRepo, presenter GetByIdCategoryPresenter) GetByIdCategoryUseCase {
	return &getByIdCategoryInteractor{repo: repo, presenter: presenter}
}

func (uc getByIdCategoryInteractor) Execute(ctx context.Context, input GetByIdCategoryInput) (GetByIdCategoryOutput, error) {
	category, err := uc.repo.GetByID(ctx, input.Id)
	if err != nil {
		return GetByIdCategoryOutput{}, err
	}
	return uc.presenter.Output(category), nil
}
