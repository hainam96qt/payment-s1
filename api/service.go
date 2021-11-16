package api

import (
	"context"
	"fmt"
	"payment-s1/configs"
	"payment-s1/pkg/entities"
	"payment-s1/pkg/usecase/service"

	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ PaymentServiceServer = &PaymentServer{}

type PaymentServer struct {
	PaymentService *service.PaymentService
}

func (s *PaymentServer) GetWager(ctx context.Context, request *GetWagerRequest) (*GetWagerResponse, error) {
	fmt.Println(request.Limit)
	var r = &entities.GetListWagerRequest{
		Limit:  request.Limit,
		Paging: request.Page,
	}
	result, err := s.PaymentService.GetListWager(ctx, r)
	if err != nil {

		return nil, err
	}
	var wagers []*Wager
	for _, v := range result.List {
		wagers = append(wagers, &Wager{
			Id:                  v.Id,
			TotalWagerValue:     v.TotalWagerValue,
			Odds:                v.Odds,
			SellingPercentage:   v.SellingPercentage,
			SellingPrice:        v.SellingPrice,
			CurrentSellingPrice: v.CurrentSellingPrice,
			PercentageSold:      v.PercentageSold,
			AmountSold:          v.AmountSold,
			PlacedAt:            timestamppb.New(v.PlacedAt),
		})
	}
	return &GetWagerResponse{
		Wager: wagers,
	}, nil
}

func (s *PaymentServer) BuyWager(ctx context.Context, request *BuyWagerRequest) (*BuyWagerResponse, error) {
	var r = &entities.BuyWagerRequest{
		WagerID:     int32(request.WagerId),
		BuyingPrice: request.BuyingPrice,
	}
	result, err := s.PaymentService.BuyWager(ctx, r)
	if err != nil {

		return nil, err
	}
	return &BuyWagerResponse{
		Id:          result.ID,
		WagerId:     result.WagerID,
		BuyingPrice: result.BuyingPrice,
		BoughtAt:    timestamppb.New(result.BoughtAt),
	}, nil
}

func (s *PaymentServer) CreateWager(ctx context.Context, request *CreateWagerRequest) (*CreateWagerResponse, error) {
	var r = &entities.CreateWagerRequest{
		TotalWagerValue:   request.TotalWagerValue,
		Odds:              request.Odds,
		SellingPercentage: float64(request.Odds),
		SellingPrice:      float64(request.SellingPrice),
	}
	result, err := s.PaymentService.CreateWager(ctx, r)
	if err != nil {

		return nil, err
	}
	return &CreateWagerResponse{
		Id:                  result.Id,
		TotalWagerValue:     result.TotalWagerValue,
		Odds:                result.Odds,
		SellingPercentage:   result.SellingPercentage,
		SellingPrice:        result.SellingPrice,
		CurrentSellingPrice: result.CurrentSellingPrice,
		PercentageSold:      result.PercentageSold,
		AmountSold:          result.AmountSold,
		PlacedAt:            timestamppb.New(result.PlacedAt),
	}, nil
}

func (s *PaymentServer) mustEmbedUnimplementedPaymentServiceServer() {
	panic("implement me")
}

func NewPaymentServer(cfg *configs.Config) (*PaymentServer, error) {
	paymentService, err := service.NewPaymentService(cfg)
	if err != nil {
		return nil, err
	}
	return &PaymentServer{
		PaymentService: paymentService,
	}, nil
}
