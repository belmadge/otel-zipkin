package main

import (
	"github.com/belmadge/otel-zipkin/service-b/handlers"
	"github.com/belmadge/otel-zipkin/service-b/tracing"

	"github.com/gin-gonic/gin"
)

func main() {
	tracing.InitTracer()

	r := gin.Default()
	r.GET("/weather/:cep", handlers.HandleWeather)
	r.Run(":8081")
}
