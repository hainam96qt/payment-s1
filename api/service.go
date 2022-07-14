package api

import (
	"context"
	"fmt"
	"payment-s1/configs"
	"payment-s1/pkg/entities"
	"payment-s1/pkg/usecase/service"

	"google.golang.org/protobuf/types/known/timestamppb"
)


// order product service
var _ OrderProductServiceServer = &OrderProductServer{}

type OrderProductServer struct {
	OrderProductService *service.OrderProductService
}

func NewOrderProductServer(cfg *configs.Config) (*OrderProductServer, error) {
	paymentService, err := service.NewOrderProductService(cfg)
	if err != nil {
		return nil, err
	}
	return &OrderProductServer{
		OrderProductService: paymentService,
	}, nil
}

func (s *OrderProductServer) Login(ctx context.Context, request *LoginRequest) (*LoginResponse, error) {
	panic("implement me")
}

func (s *OrderProductServer) GetWager(ctx context.Context, request *GetWagerRequest) (*GetWagerResponse, error) {
	fmt.Println(request.Limit)
	var r = &entities.GetListWagerRequest{
		Limit:  request.Limit,
		Paging: request.Page,
	}
	result, err := s.OrderProductService.GetListWager(ctx, r)
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

func (s *OrderProductServer) BuyWager(ctx context.Context, request *BuyWagerRequest) (*BuyWagerResponse, error) {
	var r = &entities.BuyWagerRequest{
		WagerID:     int32(request.WagerId),
		BuyingPrice: request.BuyingPrice,
	}
	result, err := s.OrderProductService.BuyWager(ctx, r)
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

func (s *OrderProductServer) CreateWager(ctx context.Context, request *CreateWagerRequest) (*CreateWagerResponse, error) {
	var r = &entities.CreateWagerRequest{
		TotalWagerValue:   request.TotalWagerValue,
		Odds:              request.Odds,
		SellingPercentage: float64(request.Odds),
		SellingPrice:      float64(request.SellingPrice),
	}
	result, err := s.OrderProductService.CreateWager(ctx, r)
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

func (s *OrderProductServer) mustEmbedUnimplementedOrderProductServiceServer() {
	panic("implement me")
}

// authentication service
var _ AuthenticationServiceServer = &AuthenticationServer{}

type AuthenticationServer struct {
	AuthenticationService *service.AuthenticationService
}

func NewAuthenticationServer(cfg *configs.Config) (*AuthenticationServer, error) {
	AuthenticationService, err := service.NewAuthenticationService(cfg)
	if err != nil {
		return nil, err
	}
	return &AuthenticationServer{
		AuthenticationService: AuthenticationService,
	}, nil
}


func (a AuthenticationServer) Login(ctx context.Context, request *LoginRequest) (*LoginResponse, error) {
	var r = &entities.LoginRequest{
		Username: request.UserName,
		Password: request.Password,
	}
	result, err := a.AuthenticationService.Login(ctx, r)
	if err != nil {
		return nil, err
	}
	return &LoginResponse{
		Token:         result.Token,
	}, nil
}

func (a AuthenticationServer) mustEmbedUnimplementedAuthenticationServiceServer() {
	panic("implement me")
}


