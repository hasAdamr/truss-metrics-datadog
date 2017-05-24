package middlewares

import (
	pb "github.com/hasAdamr/truss-metrics-datadog/metrics-service"
)

func WrapService(in pb.MetricsServer) pb.MetricsServer {
	return in
}
