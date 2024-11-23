package repo

import (
	"context"
	"neuro-most/category-service/internal/entities"
)

type categoryGORM struct {
	Id   int64  `gorm:"primaryKey"`
	Name string `gorm:"not null"`
}

type CategoryRepo struct {
	db GSQL
}

func NewCategoryRepo(db GSQL) CategoryRepo {
	db.AutoMigrate(&categoryGORM{})
	return CategoryRepo{db: db}
}

func (r CategoryRepo) Create(ctx context.Context, category entities.Category) error {
	var categoryGORM categoryGORM
	categoryGORM.Name = category.Name()
	if err := r.db.Create(ctx, &categoryGORM); err != nil {
		return err
	}
	return nil
}

func (r CategoryRepo) Update(ctx context.Context, category entities.Category) error {
	updates := map[string]interface{}{
		"name": category.Name(),
	}
	if err := r.db.UpdateOne(ctx, &updates, &categoryGORM{Id: category.ID()}, &categoryGORM{}); err != nil {
		return err
	}
	return nil
}

func (r CategoryRepo) Delete(ctx context.Context, category entities.Category) error {
	if err := r.db.Delete(ctx, &categoryGORM{}, &categoryGORM{Id: category.ID()}); err != nil {
		return err
	}
	return nil
}

func (r CategoryRepo) Fetch(ctx context.Context, page, pageSize int64) ([]entities.Category, int64, error) {
	var Category []categoryGORM
	query := r.db.BeginFind(ctx, &Category)
	var total int64
	query.Count(&total)
	query = query.Page(int(page), int(pageSize)).OrderBy("id desc")
	err := query.Find(&Category)
	if err != nil {
		return nil, 0, entities.ErrorCategoryFetch
	}
	var result []entities.Category
	for _, tag := range Category {
		result = append(result, r.convertToTag(tag))
	}
	return result, total, nil
}

func (r CategoryRepo) GetByID(ctx context.Context, id int64) (entities.Category, error) {
	var category categoryGORM
	if err := r.db.BeginFind(ctx, &category).Where(&categoryGORM{Id: id}).First(&category); err != nil {
		return entities.Category{}, entities.ErrCategoryNotFound
	}
	return r.convertToTag(category), nil
}

func (r CategoryRepo) convertToTag(CategoryGORM categoryGORM) entities.Category {
	return entities.NewCategory(
		CategoryGORM.Id,
		CategoryGORM.Name,
	)
}
