package application

import (
	"context"
	"github.com/igoramorim/go-practice-clean-arch/internal/domain/dorder"
	"github.com/igoramorim/go-practice-clean-arch/pkg/ddd"
	"log"
	"time"
)

// TODO: Add unit test.

func NewCreateOrderService(repo dorder.Repository, eventDispatcher ddd.EventDispatcher) *CreateOrderService {
	return &CreateOrderService{
		repo:            repo,
		eventDispatcher: eventDispatcher,
	}
}

var _ dorder.CreateOrderUseCase = (*CreateOrderService)(nil)

type CreateOrderService struct {
	repo            dorder.Repository
	eventDispatcher ddd.EventDispatcher
}

func (s *CreateOrderService) Execute(ctx context.Context,
	input dorder.CreateOrderUseCaseInput) (dorder.CreateOrderUseCaseOutput, error) {

	log.Printf("creating order with input: %+v\n", input)

	order, err := dorder.New(input.ID, input.Price, input.Tax)
	if err != nil {
		return dorder.CreateOrderUseCaseOutput{}, err
	}

	// TODO: Add unit of work and repo rollback if publishes fails.

	err = s.repo.Create(ctx, order)
	if err != nil {
		return dorder.CreateOrderUseCaseOutput{}, err
	}

	err = s.eventDispatcher.Dispatch(ctx, order)
	if err != nil {
		return dorder.CreateOrderUseCaseOutput{}, err
	}

	return dorder.CreateOrderUseCaseOutput{
		ID:         order.ID(),
		Price:      order.Price(),
		Tax:        order.Tax(),
		FinalPrice: order.FinalPrice(),
		CreatedAt:  order.CreatedAt().Format(time.RFC3339Nano),
	}, nil
}
