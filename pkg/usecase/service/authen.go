package service

import (
	"context"
	"database/sql"
	"github.com/dgrijalva/jwt-go"
	"payment-s1/configs"
	"payment-s1/pkg/db/mysql_db"
	"payment-s1/pkg/entities"
	sql_modle "payment-s1/pkg/usecase/model"
	"time"
)

var jwtKey string
var _ entities.Authentication = &AuthenticationService{}

type AuthenticationService struct {
	DatabaseConn *sql.DB
	Query        *sql_modle.Queries
}

func NewAuthenticationService(cfg *configs.Config) (*AuthenticationService, error) {
	jwtKey = cfg.SecretKey
	databaseConn, err := mysql_db.ConnectDatabase(cfg.Mysqldb)
	if err != nil {
		return nil, err
	}
	query := sql_modle.New(databaseConn)
	return &AuthenticationService{
		DatabaseConn: databaseConn,
		Query:        query,
	}, nil
}

type Claims struct {
	Username string `json:"username"`
	USerID int `json:"user_id"`
	jwt.StandardClaims
}


func (a AuthenticationService) Login(ctx context.Context, request *entities.LoginRequest) (*entities.LoginResponse, error) {
	// TODO validate user
	userID := 1

	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	expirationTime := time.Now().Add(5 * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Username: request.Username,
		USerID: userID,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return nil, err
	}
	return &entities.LoginResponse{Token: tokenString}, nil
}


