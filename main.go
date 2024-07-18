package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

var (
	requestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "myapp_requests_total",
			Help: "Total number of requests received.",
		},
		[]string{"uri"},
	)
)

func init() {
	// Register metrics with Prometheus client library
	prometheus.MustRegister(requestsTotal)
}

func main() {
	// Create a Gin router
	r := gin.Default()

	// 记录HTTP请求的中间件
	r.Use(func(c *gin.Context) {
		path := c.Request.URL.Path
		requestsTotal.WithLabelValues(path).Inc()
		c.Next()
	})

	// Endpoint handler for metrics scraping
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Start the HTTP server
	if err := http.ListenAndServe(":8182", r); err != nil {
		panic(err)
	}
}
