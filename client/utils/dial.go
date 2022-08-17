package utils

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Create gRPC connection
func DialGrpc() (*grpc.ClientConn, error) {
	return grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
}
