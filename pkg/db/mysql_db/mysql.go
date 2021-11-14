package mysql_db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type DatabaseConfig struct {
	Conn string
}

func ConnectDatabase(args DatabaseConfig) (*sql.DB, error) {
	db, err := sql.Open("mysql_db", args.Conn)
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}
	return db, nil
}
