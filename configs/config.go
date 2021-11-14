package configs

import (
	"challenge/pkg/db/mysql_db"
	"challenge/pkg/entities/common"
	error_code "challenge/pkg/enum/error"
	"os"
)

type Config struct {
	Host     Host
	Mysqldb  mysql_db.DatabaseConfig
}

type Host struct {
	ApiHost        string
	ApiHostToken   string
	ControllerHost string
}

const (
	EnvRabbitmqName = "AMQP_URL"
	EnvMysqldbName  = "MYSQL_URL"

	EnvApiHost      = "PLAYLIST_API_HOST"
	EnvApiHostToken = "PLAYLIST_API_TOKEN"

	EnvControllerHttp = "CAS_CONTROLLER_HOST"
)

// NewConfig returns a new decoded Config struct
func NewConfig() (*Config, error) {
	config := &Config{}

	// Database mysql_db config
	// {user}:{password}@tcp({host}:{post})/{database}?charset=utf8&parseTime=True&loc=Local
	connMysqlString := os.Getenv(EnvMysqldbName)
	if connMysqlString == "" {
		return nil, common.NewErr(error_code.Internal, "Empty string connection to mysql_db")
	}
	config.Mysqldb = mysql_db.DatabaseConfig{Conn: connMysqlString}

	// Host
	hostInfo := os.Getenv(EnvApiHost)
	if hostInfo == "" {
		return nil, common.NewErr(error_code.Internal, "Empty host info")
	}
	hostInfoToken := os.Getenv(EnvApiHostToken)
	if hostInfoToken == "" {
		return nil, common.NewErr(error_code.Internal, "Empty host info token")
	}
	// Controller host
	host := Host{
		ApiHost:        hostInfo,
		ApiHostToken:   hostInfoToken,
		ControllerHost: os.Getenv(EnvControllerHttp),
	}
	config.Host = host

	return config, nil
}
