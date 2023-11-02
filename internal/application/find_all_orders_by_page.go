package application

import (
	"context"
	"github.com/igoramorim/go-practice-clean-arch/internal/domain/dorder"
)

// TODO: Add unit test.

func NewFindAllOrdersByPageService(repo dorder.Repository) *FindAllOrdersByPageService {
	return &FindAllOrdersByPageService{
		repo: repo,
	}
}

var _ dorder.FindAllOrdersByPageUseCase = (*FindAllOrdersByPageService)(nil)

type FindAllOrdersByPageService struct {
	repo dorder.Repository
}

func (s *FindAllOrdersByPageService) Execute(ctx context.Context,
	input dorder.FindAllOrdersByPageUseCaseInput) (dorder.FindAllOrdersByPageUseCaseOutput, error) {

	orders, total, err := s.repo.FindAllByPage(ctx, input.Page, input.Limit, input.Sort)
	if err != nil {
		return dorder.FindAllOrdersByPageUseCaseOutput{}, err
	}

	output := dorder.FindAllOrdersByPageUseCaseOutput{
		Paging: dorder.Paging{
			Limit: int64(input.Limit),
			Total: total,
		},
		Orders: make([]dorder.FindAllOrdersByPageUseCaseOutputItem, input.Limit),
	}
	for i, order := range orders {
		output.Orders[i].Map(order)
	}

	return output, nil
}
