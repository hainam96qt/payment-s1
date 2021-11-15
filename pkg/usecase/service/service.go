package service

import (
	"challenge/configs"
	"challenge/pkg/db/mysql_db"
	"challenge/pkg/entities"
	sql_modle "challenge/pkg/usecase/model"
	"context"
	"database/sql"
)

var _ entities.Payment = &PaymentService{}

type PaymentService struct {
	DatabaseConn *sql.DB
	Host         configs.Host
	Query        *sql_modle.Queries
}

func NewPaymentService(cfg *configs.Config) (*PaymentService, error) {
	databaseConn, err := mysql_db.ConnectDatabase(cfg.Mysqldb)
	if err != nil {
		return nil, err
	}
	query := sql_modle.New(databaseConn)
	return &PaymentService{
		Host:         cfg.Host,
		DatabaseConn: databaseConn,
		Query:        query,
	}, nil
}

func (p2 PaymentService) HelloWorld(ctx context.Context, p entities.HelloWorldRequest) (entities.HelloWorldResponse, error) {
	return entities.HelloWorldResponse{
		Id:      p.Id,
		Message: "SomeOne",
	}, nil
}


