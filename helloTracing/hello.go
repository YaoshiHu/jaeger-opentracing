package main

import (
	"fmt"
	"github.com/opentracing/opentracing-go/log"
	"github.com/YaoshiHu/jaeger-opentracing/initTracer"
)

func main() {
	sampleTracer, closer := initTracer.SimpleTracer("Hello World")
	defer closer.Close()

	span := sampleTracer.StartSpan("hello to world")
	defer span.Finish()

	span.LogFields(
		log.String("event", "Hello Begin"),
	)

	fmt.Println("Hello World")

	span.LogFields(
		log.String("event", "Hello End"),
	)

}
