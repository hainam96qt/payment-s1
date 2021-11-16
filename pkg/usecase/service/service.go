package service

import (
	"context"
	"database/sql"
	"payment-s1/configs"
	"payment-s1/pkg/db/mysql_db"
	"payment-s1/pkg/entities"
	"payment-s1/pkg/usecase/convert"
	sql_modle "payment-s1/pkg/usecase/model"
	"time"

	"payment-s1/pkg/entities/errors"
)

var _ entities.Payment = &PaymentService{}

type PaymentService struct {
	DatabaseConn *sql.DB
	Query        *sql_modle.Queries
}

func NewPaymentService(cfg *configs.Config) (*PaymentService, error) {
	databaseConn, err := mysql_db.ConnectDatabase(cfg.Mysqldb)
	if err != nil {
		return nil, err
	}
	query := sql_modle.New(databaseConn)
	return &PaymentService{
		DatabaseConn: databaseConn,
		Query:        query,
	}, nil
}

func (p PaymentService) CreateWager(ctx context.Context, request *entities.CreateWagerRequest) (*entities.CreateWagerResponse, error) {
	if request.TotalWagerValue == 0 {
		return nil, errors.NewErr(errors.ErrorDescription, "ERROR_DESCRIPTION")
	}
	if float64(request.SellingPrice) > float64(request.TotalWagerValue)*request.SellingPercentage/100 {
		return nil, errors.NewErr(errors.ErrorDescription, "ERROR_DESCRIPTION")
	}
	if request.SellingPercentage < 0 || request.SellingPercentage > 100 {
		return nil, errors.NewErr(errors.ErrorDescription, "ERROR_DESCRIPTION")
	}
	var wagerDB = sql_modle.CreateWagerParams{
		Odds:                convert.NewNullInt64(request.Odds),
		TotalWagerValue:     convert.NewNullInt64(request.TotalWagerValue),
		SellingPercentage:   convert.NewNullInt64(request.TotalWagerValue),
		SellingPrice:        convert.NewNullFloat64(convert.ToFixed(request.SellingPrice, 2)),
		CurrentSellingPrice: convert.NewNullFloat64(request.SellingPrice),
		PlacedAt:            convert.NewNullTime(time.Now()),
	}

	err := p.Query.CreateWager(ctx, wagerDB)
	if err != nil {
		return nil, err
	}
	resultDb, err := p.Query.GetWager(ctx, sql_modle.GetWagerParams{
		Odds:            convert.NewNullInt64(request.Odds),
		TotalWagerValue: convert.NewNullInt64(request.TotalWagerValue),
	})
	if err != nil {
		return nil, err
	}
	return &entities.CreateWagerResponse{
		Id:                  resultDb.ID,
		TotalWagerValue:     resultDb.TotalWagerValue.Int64,
		Odds:                resultDb.Odds.Int64,
		SellingPercentage:   resultDb.SellingPercentage.Int64,
		SellingPrice:        resultDb.SellingPrice.Float64,
		CurrentSellingPrice: resultDb.CurrentSellingPrice.Float64,
		PercentageSold:      resultDb.PercentageSold.Float64,
		AmountSold:          resultDb.AmountSold.Int64,
		PlacedAt:            resultDb.PlacedAt.Time,
	}, nil
}

func (p PaymentService) BuyWager(ctx context.Context, request *entities.BuyWagerRequest) (*entities.BuyWagerResponse, error) {
	wagerDb, err := p.Query.GetWagerById(ctx, request.WagerID)
	if err != nil {
		return nil, err
	}

	if wagerDb.CurrentSellingPrice.Float64 < request.BuyingPrice {
		return nil, errors.NewErr(errors.ErrorDescription, "ERROR_DESCRIPTION")
	}

	// create buy wager
	var wagerDB = sql_modle.CreateBuyWagerLogParams{
		WagerID:     convert.NewNullInt32(request.WagerID),
		BuyingPrice: convert.NewNullFloat64(request.BuyingPrice),
		BoughtAt:    convert.NewNullTime(time.Now()),
	}

	err = p.Query.CreateBuyWagerLog(ctx, wagerDB)
	if err != nil {
		return nil, err
	}
	resultDb, err := p.Query.GetBuyWager(ctx, convert.NewNullInt32(request.WagerID))
	if err != nil {
		return nil, err
	}

	err = p.Query.UpdateWager(ctx, sql_modle.UpdateWagerParams{
		CurrentSellingPrice: convert.NewNullFloat64(request.BuyingPrice),
		PercentageSold:      convert.NewNullFloat64((request.BuyingPrice / float64(wagerDb.TotalWagerValue.Int64) * 100)),
		AmountSold:          convert.NewNullInt64(wagerDb.AmountSold.Int64 + 1),
		ID:                  request.WagerID,
	})
	if err != nil {
		return nil, err
	}
	return &entities.BuyWagerResponse{
		ID:          resultDb.ID,
		WagerID:     resultDb.WagerID.Int32,
		BuyingPrice: resultDb.BuyingPrice.Float64,
		BoughtAt:    resultDb.BoughtAt.Time,
	}, nil
}

func (p PaymentService) GetListWager(ctx context.Context, request *entities.GetListWagerRequest) (*entities.GetListWagerResponse, error) {
	if request.Paging <= 0 || request.Limit <= 0 {
		return nil, errors.NewErr(errors.ErrorDescription, "ERROR_DESCRIPTION1")
	}
	wagerDbs, err := p.Query.ListWagers(ctx, sql_modle.ListWagersParams{
		Limit:  request.Limit,
		Offset: (request.Paging - 1) * request.Limit,
	})
	if err != nil {
		return nil, err
	}
	var list []entities.Wager
	for _, v := range wagerDbs {
		list = append(list, entities.Wager{
			Id:                  v.ID,
			TotalWagerValue:     v.TotalWagerValue.Int64,
			Odds:                v.Odds.Int64,
			SellingPercentage:   v.SellingPercentage.Int64,
			SellingPrice:        v.SellingPrice.Float64,
			CurrentSellingPrice: v.CurrentSellingPrice.Float64,
			PercentageSold:      v.PercentageSold.Float64,
			AmountSold:          v.AmountSold.Int64,
			PlacedAt:            v.PlacedAt.Time,
		})
	}
	return &entities.GetListWagerResponse{List: list}, err
}
