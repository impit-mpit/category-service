package entities

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrCategoryNotFound = status.New(codes.NotFound, "category not found").Err()
	ErrorCategoryCreate = status.New(codes.Internal, "error create category").Err()
	ErrorCategoryUpdate = status.New(codes.Internal, "error update category").Err()
	ErrorCategoryDelete = status.New(codes.Internal, "error delete category").Err()
	ErrorCategoryFetch  = status.New(codes.Internal, "error fetch category").Err()
)

type (
	CategoryRepo interface {
		Create(ctx context.Context, tag Category) error
		Update(ctx context.Context, tag Category) error
		Delete(ctx context.Context, tag Category) error
		GetByID(ctx context.Context, id int64) (Category, error)
		Fetch(ctx context.Context, page, pageSize int64) ([]Category, int64, error)
	}

	Category struct {
		id   int64
		name string
	}
)

func NewCategory(
	id int64,
	name string,
) Category {
	return Category{
		id:   id,
		name: name,
	}
}

func NewCategoryCreate(
	name string,
) Category {
	return Category{
		name: name,
	}
}

func (t Category) ID() int64 {
	return t.id
}

func (t Category) Name() string {
	return t.name
}

func (t *Category) SetID(id int64) {
	t.id = id
}

func (t *Category) SetName(name string) {
	t.name = name
}
