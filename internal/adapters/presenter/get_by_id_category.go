package presenter

import (
	"neuro-most/category-service/internal/entities"
	"neuro-most/category-service/internal/usecase"
)

type GetByIdPresenter struct {
}

func NewGetByIdPresenter() GetByIdPresenter {
	return GetByIdPresenter{}
}

func (p GetByIdPresenter) Output(category entities.Category) usecase.GetByIdCategoryOutput {
	return usecase.GetByIdCategoryOutput{
		Id:   category.ID(),
		Name: category.Name(),
	}
}
