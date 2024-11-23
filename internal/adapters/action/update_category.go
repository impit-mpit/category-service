package action

import (
	"context"
	categoryv1 "neuro-most/category-service/gen/go/category/v1"
	"neuro-most/category-service/internal/usecase"
)

type UpdateCategoryAction struct {
	uc usecase.UpdateCategoryUseCase
}

func NewUpdateCategoryAction(uc usecase.UpdateCategoryUseCase) UpdateCategoryAction {
	return UpdateCategoryAction{uc: uc}
}

func (a UpdateCategoryAction) Execute(ctx context.Context, input *categoryv1.UpdateCategoryRequest) error {
	return a.uc.Execute(ctx, usecase.UpdateCategoryInput{Id: input.Id, Name: input.Name})
}
