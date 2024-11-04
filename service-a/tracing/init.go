package tracing

import (
	"context"
	"log"
	"net/http"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/zipkin"
	"go.opentelemetry.io/otel/propagation"
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

func AddTrace(req *http.Request) {
	ctx, span := Tracer.Start(context.Background(), "forward_to_service_b")
	defer span.End()

	otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(req.Header))
}
