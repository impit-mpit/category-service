package action

import (
	"context"
	categoryv1 "neuro-most/category-service/gen/go/category/v1"
	"neuro-most/category-service/internal/usecase"
)

type GetByIDCategoryAction struct {
	uc usecase.GetByIdCategoryUseCase
}

func NewGetByIDCategoryAction(uc usecase.GetByIdCategoryUseCase) GetByIDCategoryAction {
	return GetByIDCategoryAction{uc: uc}
}

func (a GetByIDCategoryAction) Execute(ctx context.Context, input *categoryv1.GetCategoryByIdRequest) (*categoryv1.Category, error) {
	category, err := a.uc.Execute(ctx, usecase.GetByIdCategoryInput{Id: input.Id})
	if err != nil {
		return nil, err
	}
	return &categoryv1.Category{
		Id:   category.Id,
		Name: category.Name,
	}, nil
}
