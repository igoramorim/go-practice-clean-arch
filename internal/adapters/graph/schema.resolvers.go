package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"context"
	"github.com/pkg/errors"
	"log"

	"github.com/igoramorim/go-practice-clean-arch/internal/adapters/graph/model"
	"github.com/igoramorim/go-practice-clean-arch/internal/domain/dorder"
)

// CreateOrder is the resolver for the createOrder field.
func (r *mutationResolver) CreateOrder(ctx context.Context, input *model.CreateOrderInput) (*model.Order, error) {
	in := dorder.CreateOrderUseCaseInput{
		ID:    int64(input.ID),
		Price: input.Price,
		Tax:   input.Tax,
	}
	out, err := r.CreateOrderUseCase.Execute(ctx, in)
	if err != nil {
		log.Printf("[ERROR] %s\n", errors.WithMessage(err, "graphql creating order").Error())
		return nil, err
	}

	return model.MapCreateOrderOutput(out), nil
}

// FindAllOrdersByPage is the resolver for the findAllOrdersByPage field.
func (r *queryResolver) FindAllOrdersByPage(ctx context.Context, input *model.FindAllOrdersByPageInput) (*model.FindAllOrdersByPageOutput, error) {
	var (
		page, limit int
		sort        string
	)
	if input.Page == nil {
		page = 1
	}
	if input.Limit == nil {
		limit = 10
	}
	in := dorder.FindAllOrdersByPageUseCaseInput{
		Page:  page,
		Limit: limit,
		Sort:  sort,
	}
	out, err := r.FindAllOrdersByPageUseCase.Execute(ctx, in)
	if err != nil {
		log.Printf("[ERROR] %s\n", errors.WithMessage(err, "graphql listing orders").Error())
		return nil, err
	}

	return model.MapFindAllOrdersByPageOutput(out), nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
