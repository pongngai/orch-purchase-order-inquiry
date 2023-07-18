package main

import (
	"com.ai.orch-purchase-order-inquiry/grpc/server"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	// Create a new gRPC server
	srv := grpc.NewServer()

	// Register the gRPC service implementation
	server.RegisterOrderServiceServer(srv)

	// Start the gRPC server
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Println("gRPC server listening on :50051")
	if err := srv.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
