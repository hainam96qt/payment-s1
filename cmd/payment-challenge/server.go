package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"payment-s1/api"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func RunServerGRPC(ctx context.Context, v1API *api.OrderProductServer, authen *api.AuthenticationServer, port string) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	// register service
	server := grpc.NewServer()
	api.RegisterOrderProductServiceServer(server, v1API)
	api.RegisterAuthenticationServiceServer(server, authen)

	// start gRPC server
	log.Println("starting gRPC server at port:" + port)
	return server.Serve(listen)
}

func RunServerREST(ctx context.Context, gPRCPort, httpPort string) error {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := api.RegisterOrderProductServiceHandlerFromEndpoint(ctx, mux, ":"+gPRCPort, opts)
	if err != nil {
		return err
	}
	// start gRPC server
	log.Println("starting REST server at port:" + httpPort)
	return http.ListenAndServe(":"+httpPort, mux)
}
