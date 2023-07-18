package server

import (
	"com.ai.orch-purchase-order-inquiry/infrastructure/gormrepository"
	pb "com.ai.orch-purchase-order-inquiry/proto"
	"context"
	"google.golang.org/grpc"
	"strconv"
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
	orderId, _ := strconv.Atoi(req.Id)
	order, err := gormrepository.GetByOrderId(orderId)
	if err != nil {
		return nil, err
	}
	return &pb.GetOrderResponse{
		OrderId:     int64(order.ItemId),
		Price:       int64(order.Price),
		Quantity:    int64(order.Quantity),
		CustomerId:  int64(order.CustomerId),
		TotalAmount: int64(order.TotalAmount),
		Status:      string(order.Status),
	}, nil
}
