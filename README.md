# truss-metrics-datadog

This is a very simple [truss](https://github.com/TuneLab/go-truss) service that
uses LabeledMiddlewares to log metrics to datadog.

Checkout ./metrics-service/middlewares/endpoints.go

This does not include configuring dogstatsd.
