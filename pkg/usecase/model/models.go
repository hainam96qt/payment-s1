// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"database/sql"
)

type Payment struct {
	ID        string         `json:"id"`
	RequestID sql.NullString `json:"request_id"`
}
