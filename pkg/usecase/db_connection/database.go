package db_connection

import (
	"challenge/configs"
	"database/sql"
	"challenge/pkg/entities"
	sql_model "challenge/pkg/usecase/model"
	"challenge/pkg/db/mysql_db"
)

var _ entities.Payment = &PaymentService{}

type PaymentService struct {
	DatabaseConn *sql.DB
	Host         configs.Host
	Query        *sql_model.Queries
}

func (p PaymentService) Init() {
	panic("implement me")
}

func NewTaskService(cfg *configs.Config) (*PaymentService, error) {
	databaseConn, err := mysql_db.ConnectDatabase(cfg.Mysqldb)
	if err != nil {
		return nil, err
	}
	query := sql_model.New(databaseConn)
	return &PaymentService{
		Host:         cfg.Host,
		DatabaseConn: databaseConn,
		Query:        query,
	}, nil
}
