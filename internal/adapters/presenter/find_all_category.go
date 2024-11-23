package presenter

import (
	"neuro-most/category-service/internal/entities"
	"neuro-most/category-service/internal/usecase"
)

type FindAllCategoryPresenter struct {
}

func NewFindAllCategoryPresenter() FindAllCategoryPresenter {
	return FindAllCategoryPresenter{}
}

func (p FindAllCategoryPresenter) Output(categories []entities.Category) []usecase.FindAllCategoryOutput {
	var res []usecase.FindAllCategoryOutput
	for _, category := range categories {
		res = append(res, usecase.FindAllCategoryOutput{
			Id:   category.ID(),
			Name: category.Name(),
		})
	}
	return res
}
