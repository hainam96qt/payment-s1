package api

import (
	"challenge/configs"
	"challenge/pkg/entities"
	"challenge/pkg/usecase/service"
	"context"
)

var _ PaymentServiceServer = &PaymentServer{}
type PaymentServer struct {
	PaymentService *service.PaymentService
}

func (s *PaymentServer) mustEmbedUnimplementedPaymentServiceServer() {
	panic("implement me")
}

func NewPaymentServer(cfg *configs.Config) (*PaymentServer, error) {
	paymentService, err := service.NewTaskService(cfg)
	if err != nil {
		return nil, err
	}
	return &PaymentServer{
		PaymentService: paymentService,
	}, nil
}

func (s *PaymentServer) HelloWorld(ctx context.Context, req *HelloWorldRequest) (*HelloWorldResponse, error) {
	var r= &entities.HelloWorldRequest{
		Id: int(req.Id),
	}
	result, err := s.PaymentService.HelloWorld(ctx, *r)
	if err != nil {

		return nil, err
	}
	return &HelloWorldResponse{
		RequestId:     int64(result.Id),
		Message:       "",
	}, nil
}
