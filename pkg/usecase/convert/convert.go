package convert

import (
	"database/sql"
	"math"
	"time"
)

func NewNullInt32(arg int32) sql.NullInt32 {
	return sql.NullInt32{
		Int32: arg,
		Valid: true,
	}
}

func NewNullInt64(arg int64) sql.NullInt64 {
	return sql.NullInt64{
		Int64: arg,
		Valid: true,
	}
}
func NewNullFloat64(arg float64) sql.NullFloat64 {
	return sql.NullFloat64{
		Float64: float64(arg),
		Valid:   true,
	}
}

func NewNullTime(t time.Time) sql.NullTime {
	return sql.NullTime{
		Time:  t,
		Valid: true,
	}
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func ToFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}
