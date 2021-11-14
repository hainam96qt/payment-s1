package entities

import (
	"context"
)

type Payment interface {
	Init()
	HelloWorld(ctx context.Context, p HelloWorldRequest) (HelloWorldResponse, error)
}

type HelloWorldRequest struct{
	Id int
}


type HelloWorldResponse struct{
	Id int
	Message string
}
