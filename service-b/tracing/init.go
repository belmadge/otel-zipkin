package tracing

import (
	"context"
	"log"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/zipkin"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

var Tracer trace.Tracer

func InitTracer() {
	exporter, err := zipkin.New("http://localhost:9411/api/v2/spans")
	if err != nil {
		log.Fatalf("erro ao criar o exportador Zipkin: %v", err)
	}

	tp := sdktrace.NewTracerProvider(sdktrace.WithBatcher(exporter))
	otel.SetTracerProvider(tp)
	Tracer = tp.Tracer("weather-api")
}

func StartSpan(ctx context.Context, spanName string) (context.Context, trace.Span) {
	return Tracer.Start(ctx, spanName)
}
