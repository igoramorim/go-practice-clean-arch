package application

import (
	"context"
	"github.com/igoramorim/go-practice-clean-arch/internal/domain/dorder"
)

func NewFindAllOrdersByPageService(repo dorder.Repository) *FindAllOrdersByPageService {
	return &FindAllOrdersByPageService{
		repo: repo,
	}
}

var _ dorder.FindAllOrdersByPageUseCase = (*FindAllOrdersByPageService)(nil)

type FindAllOrdersByPageService struct {
	repo dorder.Repository
}

func (s *FindAllOrdersByPageService) Execute(ctx context.Context, input dorder.FindAllOrdersByPageUseCaseInput) ([]dorder.FindAllOrdersByPageUseCaseOutput, error) {
	orders, err := s.repo.FindAllByPage(ctx, input.Page, input.Limit, input.Sort)
	if err != nil {
		return nil, err
	}

	output := make([]dorder.FindAllOrdersByPageUseCaseOutput, len(orders))
	for i, order := range orders {
		output[i].Map(order)
	}

	return output, nil
}
