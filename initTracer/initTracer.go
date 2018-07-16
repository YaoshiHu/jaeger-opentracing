package initTracer

import (
	"fmt"
	"os"

	opentracing "github.com/opentracing/opentracing-go"
	jaeger "github.com/uber/jaeger-client-go"
	config "github.com/uber/jaeger-client-go/config"
)

// SimpleTracer returns an instance of Jaeger Tracer sampling all spans, same as the one in the tutorial
func SimpleTracer(service string) (opentracing.Tracer, io.Closer) {
	cfg := &config.Configuration{
		Sampler: &config.SampleConfig{
			Type: "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpan: true,
		},
	}
	tracer, closer, err := cfg.New(service, config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger Tracer: %v\n", err))
	}
	return tracer, closer
}
