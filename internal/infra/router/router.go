package router

import (
	"context"
	"log"
	"net"
	categoryv1 "neuro-most/category-service/gen/go/category/v1"
	"neuro-most/category-service/internal/adapters/action"
	"neuro-most/category-service/internal/adapters/presenter"
	"neuro-most/category-service/internal/adapters/repo"
	"neuro-most/category-service/internal/usecase"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Router struct {
	db repo.GSQL
	categoryv1.UnimplementedCategoryServiceServer
}

func NewRouter(db repo.GSQL) Router {
	return Router{db: db}
}

func (r *Router) Listen() {
	port := ":3001"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts = []grpc.ServerOption{}
	srv := grpc.NewServer(opts...)
	categoryv1.RegisterCategoryServiceServer(srv, r)

	log.Printf("Starting gRPC server on port %s\n", port)
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (r *Router) CreateCategory(ctx context.Context, input *categoryv1.CreateCategoryRequest) (*emptypb.Empty, error) {
	var (
		uc  = usecase.NewCreateCategoryInteractor(repo.NewCategoryRepo(r.db))
		act = action.NewCreateCategoryAction(uc)
	)
	return nil, act.Execute(ctx, input)
}
func (r *Router) DeleteCategory(ctx context.Context, input *categoryv1.DeleteCategoryRequest) (*emptypb.Empty, error) {
	var (
		uc  = usecase.NewDeleteCategoryInteractor(repo.NewCategoryRepo(r.db))
		act = action.NewDeleteCategoryAction(uc)
	)
	return nil, act.Execute(ctx, input)
}
func (r *Router) GetCategoryById(ctx context.Context, input *categoryv1.GetCategoryByIdRequest) (*categoryv1.Category, error) {
	var (
		uc  = usecase.NewGetByIdCategoryInteractor(repo.NewCategoryRepo(r.db), presenter.NewGetByIdPresenter())
		act = action.NewGetByIDCategoryAction(uc)
	)
	return act.Execute(ctx, input)
}
func (r *Router) GetCategoryFeed(ctx context.Context, input *categoryv1.GetCategoryFeedRequest) (*categoryv1.GetCategoryFeedResponse, error) {
	var (
		uc  = usecase.NewFindAllCategoryInteractor(repo.NewCategoryRepo(r.db), presenter.NewFindAllCategoryPresenter())
		act = action.NewFindAllCategoryAction(uc)
	)
	return act.Execute(ctx, input)
}
func (r *Router) UpdateCategory(ctx context.Context, input *categoryv1.UpdateCategoryRequest) (*emptypb.Empty, error) {
	var (
		uc  = usecase.NewUpdateCategoryInteractor(repo.NewCategoryRepo(r.db))
		act = action.NewUpdateCategoryAction(uc)
	)
	return nil, act.Execute(ctx, input)
}
