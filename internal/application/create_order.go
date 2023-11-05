package application

import (
	"context"
	"github.com/igoramorim/go-practice-clean-arch/internal/domain/dorder"
	"github.com/igoramorim/go-practice-clean-arch/pkg/ddd"
	"time"
)

// TODO: Add unit test.

func NewCreateOrderService(repo dorder.Repository, publisher ddd.Publisher) *CreateOrderService {
	return &CreateOrderService{
		repo:      repo,
		publisher: publisher,
	}
}

var _ dorder.CreateOrderUseCase = (*CreateOrderService)(nil)

type CreateOrderService struct {
	repo      dorder.Repository
	publisher ddd.Publisher
}

func (s *CreateOrderService) Execute(ctx context.Context,
	input dorder.CreateOrderUseCaseInput) (dorder.CreateOrderUseCaseOutput, error) {

	order, err := dorder.New(input.ID, input.Price, input.Tax)
	if err != nil {
		return dorder.CreateOrderUseCaseOutput{}, err
	}

	// TODO: Add unit of work and repo rollback if publishes fails.

	err = s.repo.Create(ctx, order)
	if err != nil {
		return dorder.CreateOrderUseCaseOutput{}, err
	}

	err = s.publisher.Publish(ctx, order)
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
