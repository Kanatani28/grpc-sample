package main

import (
	"context"

	pb "../pb"
)

type incrementService struct{}

func (s *incrementService) Increment(ctx context.Context, req *pb.IncrementRequest) (*pb.IncrementResponse, error) {
	n := req.GetNumber() + 1
	return &pb.IncrementResponse{Number: n}, nil
}
