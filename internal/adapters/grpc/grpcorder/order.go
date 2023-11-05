package grpcorder

import (
	"context"
	"github.com/igoramorim/go-practice-clean-arch/internal/adapters/grpc/pb"
	"github.com/igoramorim/go-practice-clean-arch/internal/domain/dorder"
)

// TODO: Add unit tests.

func NewService(
	createOrderUseCase dorder.CreateOrderUseCase,
	findAllOrdersByPageUseCase dorder.FindAllOrdersByPageUseCase) *OrderService {

	return &OrderService{
		createOrderUseCase:         createOrderUseCase,
		findAllOrdersByPageUseCase: findAllOrdersByPageUseCase,
	}
}

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	createOrderUseCase         dorder.CreateOrderUseCase
	findAllOrdersByPageUseCase dorder.FindAllOrdersByPageUseCase
}

func (s *OrderService) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	input := dorder.CreateOrderUseCaseInput{
		ID:    req.Id,
		Price: float64(req.Price),
		Tax:   float64(req.Tax),
	}
	res, err := s.createOrderUseCase.Execute(ctx, input)
	if err != nil {
		return nil, err
	}

	return &pb.CreateOrderResponse{
		Id:         res.ID,
		Price:      float32(res.Price),
		Tax:        float32(res.Tax),
		FinalPrice: float32(res.FinalPrice),
		CreatedAt:  res.CreatedAt,
	}, nil
}

func (s *OrderService) FindAllOrdersByPage(ctx context.Context, req *pb.FindAllOrdersByPageRequest) (*pb.FindAllOrdersByPageResponse, error) {
	input := dorder.FindAllOrdersByPageUseCaseInput{
		Page:  int(req.Page),
		Limit: int(req.Limit),
		Sort:  req.Sort,
	}
	res, err := s.findAllOrdersByPageUseCase.Execute(ctx, input)
	if err != nil {
		return nil, err
	}

	output := make([]*pb.FindAllOrdersByPageItem, 0, len(res.Orders))
	for _, r := range res.Orders {
		output = append(output, &pb.FindAllOrdersByPageItem{
			Id:         r.ID,
			Price:      float32(r.Price),
			Tax:        float32(r.Tax),
			FinalPrice: float32(r.FinalPrice),
			CreatedAt:  r.CreatedAt,
		})
	}

	return &pb.FindAllOrdersByPageResponse{
		Paging: &pb.Paging{
			Limit: res.Paging.Limit,
			Total: res.Paging.Total,
		},
		Orders: output,
	}, nil
}
