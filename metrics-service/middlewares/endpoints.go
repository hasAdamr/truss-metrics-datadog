package middlewares

import (
	"golang.org/x/net/context"
	"os"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/metrics/dogstatsd"

	"github.com/hasAdamr/truss-metrics-datadog/metrics-service/svc"
)

var dogStatsDAddr = "127.0.0.1:8125"

// WrapEndpoints accepts the service's entire collection of endpoints, so that a
// set of middlewares can be wrapped around every middleware (e.g., access
// logging and instrumentation), and others wrapped selectively around some
// endpoints and not others (e.g., endpoints requiring authenticated access).
// Note that the final middleware wrapped will be the outermost middleware
// (i.e. applied first)
func WrapEndpoints(in svc.Endpoints) svc.Endpoints {
	dsd := dogstatsd.New("test.metrics-service.", log.NewJSONLogger(os.Stderr))
	ticker := time.NewTicker(time.Second)
	go dsd.SendLoop(ticker.C, "udp", dogStatsDAddr)

	// Pass in LabeledMiddlewares you want applied to every endpoint.
	// These middlewares get passed the handlers name as their first argument when applied.
	// This can be used to write generic metric gathering middlewares that can
	// report the handler name for free.
	in.WrapAllLabeledExcept(requestLatency(dsd.NewHistogram("requestLatency.", 1)))
	in.WrapAllLabeledExcept(errorCounter(dsd.NewCounter("errorCount.", 1)))

	return in
}

func requestLatency(h metrics.Histogram) svc.LabeledMiddleware {
	return func(name string, in endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			defer func(begin time.Time) {
				h.With("handler", name).Observe(time.Since(begin).Seconds())
			}(time.Now())
			return in(ctx, req)
		}
	}
}

// errorCounter is a LabeledMiddleware, when applied with WrapAllLabeledExcept
// name will be populated with the handler name, and such this middleware will
// report errors to the metric provider with the handler name. Feel free to
// remove this example middleware
func errorCounter(errCount metrics.Counter) svc.LabeledMiddleware {
	return func(name string, in endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			resp, err := in(ctx, req)
			if err != nil {
				errCount.With("handler", name).Add(1)
			}
			return resp, err
		}
	}
}
