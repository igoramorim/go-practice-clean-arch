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

var _ dorder.FindAllByPageUseCase = (*FindAllOrdersByPageService)(nil)

type FindAllOrdersByPageService struct {
	repo dorder.Repository
}

func (s *FindAllOrdersByPageService) Execute(ctx context.Context, input dorder.FindAllByPageUseCaseInput) ([]dorder.FindAllByPageUseCaseOutput, error) {
	orders, err := s.repo.FindAllByPage(ctx, input.Page, input.Limit, input.Sort)
	if err != nil {
		return nil, err
	}

	output := make([]dorder.FindAllByPageUseCaseOutput, 0, len(orders))
	for i, out := range output {
		out.Map(orders[i])
	}

	return output, nil
}
