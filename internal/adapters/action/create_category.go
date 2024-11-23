package action

import (
	"context"
	categoryv1 "neuro-most/category-service/gen/go/category/v1"
	"neuro-most/category-service/internal/usecase"
)

type CreateCategoryAction struct {
	uc usecase.CreateCategoryUseCase
}

func NewCreateCategoryAction(uc usecase.CreateCategoryUseCase) CreateCategoryAction {
	return CreateCategoryAction{uc: uc}
}

func (a CreateCategoryAction) Execute(ctx context.Context, input *categoryv1.CreateCategoryRequest) error {
	return a.uc.Execute(ctx, usecase.CreateCategoryInput{Name: input.Name})
}
