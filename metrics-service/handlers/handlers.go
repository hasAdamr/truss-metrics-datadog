package handlers

import (
	"golang.org/x/net/context"

	pb "github.com/hasAdamr/truss-metrics-datadog/metrics-service"
)

// NewService returns a naïve, stateless implementation of Service.
func NewService() pb.MetricsServer {
	return metricsService{}
}

type metricsService struct{}

// Fast implements Service.
func (s metricsService) Fast(ctx context.Context, in *pb.Empty) (*pb.Empty, error) {
	var resp pb.Empty
	resp = pb.Empty{}
	return &resp, nil
}

// Slow implements Service.
func (s metricsService) Slow(ctx context.Context, in *pb.Empty) (*pb.Empty, error) {
	var resp pb.Empty
	resp = pb.Empty{}
	return &resp, nil
}

// RandomError implements Service.
func (s metricsService) RandomError(ctx context.Context, in *pb.Empty) (*pb.Empty, error) {
	var resp pb.Empty
	resp = pb.Empty{}
	return &resp, nil
}
