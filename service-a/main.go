package main

import (
	"github.com/belmadge/otel-zipkin/service-a/handlers"
	"github.com/belmadge/otel-zipkin/service-a/tracing"

	"github.com/gin-gonic/gin"
)

func main() {
	tracing.InitTracer()

	r := gin.Default()
	r.POST("/input", handlers.HandleInput)
	r.Run(":8080")
}
