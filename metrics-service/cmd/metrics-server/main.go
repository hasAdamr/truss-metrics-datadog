// Code generated by truss.
// Rerunning truss will overwrite this file.
// DO NOT EDIT!

package main

import (
	"flag"
	"os"

	// Go Kit
	"github.com/go-kit/kit/log"

	// This Service
	"github.com/hasAdamr/truss-metrics-datadog/metrics-service/svc/server"
	"github.com/hasAdamr/truss-metrics-datadog/metrics-service/svc/server/cli"
)

func main() {
	// Update addresses if they have been overwritten by flags
	flag.Parse()

	// Logging domain.
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stdout)
		logger = log.NewContext(logger).With("ts", log.DefaultTimestampUTC)
		logger = log.NewContext(logger).With("caller", log.DefaultCaller)
	}
	server.Run(cli.Config, logger)
}