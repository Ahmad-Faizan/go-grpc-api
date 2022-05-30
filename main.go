package main

import (
	"fmt"
	"log"
	"net"
	"os"

	protos "github.com/Ahmad-faizan/go-grpc-api/protos/currency"
	"github.com/Ahmad-faizan/go-grpc-api/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	// Create a logger
	l := log.New(os.Stdout, "currency-api ", log.LstdFlags)

	// Create a new grpc server
	grpcSrv := grpc.NewServer()

	// Get an instance of the currency server
	curr := server.NewCurrency(l)

	// Register the server
	protos.RegisterCurrencyServer(grpcSrv, curr)

	// register the reflection service which allows clients to determine the methods
	// for this gRPC service
	reflection.Register(grpcSrv)

	// create a TCP socket for inbound server connections
	t, err := net.Listen("tcp", fmt.Sprintf(":%d", 9092))
	if err != nil {
		l.Println("Unable to create listener", "error", err)
		os.Exit(1)
	}

	// listen for requests
	grpcSrv.Serve(t)

}
