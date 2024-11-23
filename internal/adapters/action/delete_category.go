package action

import (
	"context"
	categoryv1 "neuro-most/category-service/gen/go/category/v1"
	"neuro-most/category-service/internal/usecase"
)

type DeleteCategoryAction struct {
	uc usecase.DeleteCategoryUseCase
}

func NewDeleteCategoryAction(uc usecase.DeleteCategoryUseCase) DeleteCategoryAction {
	return DeleteCategoryAction{uc: uc}
}

func (a DeleteCategoryAction) Execute(ctx context.Context, input *categoryv1.DeleteCategoryRequest) error {
	return a.uc.Execute(ctx, usecase.DeleteCategoryInput{Id: input.Id})
}
