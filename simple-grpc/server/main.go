package main

import (
	"context"
	pb "grpc-calculator/proto"
	"log"
)

type AddRequest struct {
	A int32
	B int32
}

type AddResponse struct {
	Result int32
}

type MultiplyRequest struct {
	A int32
	B int32
}

type MultiplyResponse struct {
	Result int32
}

type CalculatorServer struct {
	pb.UnimplementedCalculatorServer
}

func (s *CalculatorServer) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
	log.Printf("Received Add request: %d + %d", req.A, req.B)

	return &pb.AddResponse{Result: req.A + req.B}, nil
}

func (s *CalculatorServer) Multiply(ctx context.Context, req *pb.MultiplyRequest) (*pb.MultiplyResponse, error) {
	log.Printf("Received Multiply request: %d * %d", req.A, req.B)
	return &pb.MultiplyResponse{Result: req.A * req.B}, nil
}
