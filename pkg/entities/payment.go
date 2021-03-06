package entities

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Authentication interface {
	Login (context.Context, *LoginRequest) (*LoginResponse, error)
}

type OrderProduct interface {
	CreateWager(context.Context, *CreateWagerRequest) (*CreateWagerResponse, error)
	BuyWager(context.Context, *BuyWagerRequest) (*BuyWagerResponse, error)
	GetListWager(context.Context, *GetListWagerRequest) (*GetListWagerResponse, error)
}

type LoginRequest struct {
	 Username string
	 Password string
}

type LoginResponse struct {
	 Token string
}

type GetListWagerRequest struct {
	Limit  int32
	Paging int32
}

type GetListWagerResponse struct {
	List []Wager
}

type BuyWagerRequest struct {
	WagerID     int32
	BuyingPrice float64
}

type BuyWagerResponse struct {
	ID          int32
	WagerID     int32
	BuyingPrice float64
	BoughtAt    time.Time
}

type CreateWagerRequest struct {
	TotalWagerValue   int64
	Odds              int64
	SellingPercentage float64
	SellingPrice      float64
}

type CreateWagerResponse struct {
	Id                  int32
	TotalWagerValue     int64
	Odds                int64
	SellingPercentage   int64
	SellingPrice        float64
	CurrentSellingPrice float64
	PercentageSold      float64
	AmountSold          int64
	PlacedAt            time.Time
}

type Wager struct {
	Id                  int32
	TotalWagerValue     int64
	Odds                int64
	SellingPercentage   int64
	SellingPrice        float64
	CurrentSellingPrice float64
	PercentageSold      float64
	AmountSold          int64
	PlacedAt            time.Time
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
