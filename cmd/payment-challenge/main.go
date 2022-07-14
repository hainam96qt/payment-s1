package main

import (
	"context"
	"log"
	"payment-s1/api"
	"payment-s1/configs"
)

/*grpc server*/
func main() {
	var ctx = context.Background()
	cfg, err := configs.NewConfig()
	if err != nil {
		log.Print(err)
	}
	v1API, err := api.NewOrderProductServer(cfg)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	authenAPI, err := api.NewAuthenticationServer(cfg)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	go func() {
		log.Fatal(RunServerGRPC(ctx, v1API, authenAPI, "8080"))
		return
	}()

	log.Fatal(RunServerREST(ctx, "8080", "8082"))

}
