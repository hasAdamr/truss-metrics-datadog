package handlers

import (
	"errors"
	"golang.org/x/net/context"
	"math/rand"
	"time"

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
	random := rand.Intn(100)
	if random > 90 {
		time.Sleep(time.Second * time.Duration(2))
	}
	return &resp, nil
}

// Slow implements Service.
func (s metricsService) Slow(ctx context.Context, in *pb.Empty) (*pb.Empty, error) {
	var resp pb.Empty
	random := rand.Intn(5) + 3
	if random != 0 {
		time.Sleep(time.Second * time.Duration(random))
	}
	return &resp, nil
}

// RandomError implements Service.
func (s metricsService) RandomError(ctx context.Context, in *pb.Empty) (*pb.Empty, error) {
	random := rand.Intn(2)
	if random == 0 {
		return &pb.Empty{}, nil
	}
	time.Sleep(time.Second * time.Duration(random))
	return nil, errors.New("sometimes things go wrong, not this time though, this is a test")
}
