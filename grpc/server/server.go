package server

import (
	"context"
	"google.golang.org/grpc"

	pb "com.ai.orch-purchase-order-inquiry/proto"
)

// OrderServiceServer is the implementation of the gRPC service.
type OrderServiceServer struct {
	pb.UnimplementedOrderServiceServer
}

// RegisterOrderServiceServer registers the OrderServiceServer implementation to the gRPC server.
func RegisterOrderServiceServer(s *grpc.Server) {
	pb.RegisterOrderServiceServer(s, &OrderServiceServer{})
}

// GetOrder retrieves a order by ID.
func (s *OrderServiceServer) GetOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.GetOrderResponse, error) {
	return &pb.GetOrderResponse{
		OrderId:     1,
		Price:       1,
		Quantity:    1,
		CustomerId:  1,
		TotalAmount: 1,
		Status:      "SUCCESS",
	}, nil
}
