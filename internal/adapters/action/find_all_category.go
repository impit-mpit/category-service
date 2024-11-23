package action

import (
	"context"
	categoryv1 "neuro-most/category-service/gen/go/category/v1"
	"neuro-most/category-service/internal/usecase"
)

type FindAllCategoryAction struct {
	uc usecase.FindAllCategory
}

func NewFindAllCategoryAction(uc usecase.FindAllCategory) FindAllCategoryAction {
	return FindAllCategoryAction{uc: uc}
}

func (a FindAllCategoryAction) Execute(ctx context.Context, input *categoryv1.GetCategoryFeedRequest) (*categoryv1.GetCategoryFeedResponse, error) {
	var usecaseInput usecase.FindAllCategoryInput
	usecaseInput.Page = int64(input.Page)
	usecaseInput.PageSize = int64(input.PageSize)
	categories, total, err := a.uc.Execute(ctx, usecaseInput)
	if err != nil {
		return nil, err
	}
	var categoriesResponse []*categoryv1.Category
	for _, Category := range categories {
		categoriesResponse = append(categoriesResponse, &categoryv1.Category{
			Id:   Category.Id,
			Name: Category.Name,
		})
	}
	return &categoryv1.GetCategoryFeedResponse{
		Category: categoriesResponse,
		Total:    int32(total),
	}, nil
}
