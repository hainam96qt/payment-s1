// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"database/sql"
)

type BuyWagerLog struct {
	ID          int32           `json:"id"`
	WagerID     sql.NullInt32   `json:"wager_id"`
	BuyingPrice sql.NullFloat64 `json:"buying_price"`
	BoughtAt    sql.NullTime    `json:"bought_at"`
}

type Wager struct {
	ID                  int32           `json:"id"`
	Odds                sql.NullInt64   `json:"odds"`
	TotalWagerValue     sql.NullInt64   `json:"total_wager_value"`
	SellingPercentage   sql.NullInt64   `json:"selling_percentage"`
	SellingPrice        sql.NullFloat64 `json:"selling_price"`
	CurrentSellingPrice sql.NullFloat64 `json:"current_selling_price"`
	PercentageSold      sql.NullFloat64 `json:"percentage_sold"`
	AmountSold          sql.NullInt64   `json:"amount_sold"`
	PlacedAt            sql.NullTime    `json:"placed_at"`
}
