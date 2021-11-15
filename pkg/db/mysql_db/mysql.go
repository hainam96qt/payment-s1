package mysql_db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type DatabaseConfig struct {
	Username string  `yaml:"username"`
	Password string `yaml:"password"`
	Port string `yaml:"post"`
	DatabaseName string `yaml:"database_name"`
}

func ConnectDatabase(args DatabaseConfig) (*sql.DB, error) {
	db, err := sql.Open("mysql", fmt.Sprint("%s:%s@tcp(127.0.0.1:%s)/%s)", args.Username, args.Password, args.Port, args.DatabaseName))
	if err != nil {
		log.Print(err.Error())
		return nil, err
	}
	return db, nil
}
