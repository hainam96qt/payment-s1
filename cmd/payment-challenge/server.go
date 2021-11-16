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

// RunServer runs gRPC service to publish ToDo service
func RunServerGRPC(ctx context.Context, v1API *api.PaymentServer, port string) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	// register service
	server := grpc.NewServer()
	api.RegisterPaymentServiceServer(server, v1API)

	// start gRPC server
	log.Println("starting gRPC server at port:" + port)
	return server.Serve(listen)
}

func RunServerREST(ctx context.Context, gPRCPort, httpPort string) error {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := api.RegisterPaymentServiceHandlerFromEndpoint(ctx, mux, ":"+gPRCPort, opts)
	if err != nil {
		return err
	}

	// start gRPC server
	log.Println("starting REST server at port:" + httpPort)
	return http.ListenAndServe(":"+httpPort, mux)
}
