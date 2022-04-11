package middleware

import (
	"goprom/metric"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

func PrometheusUriRequestTotal() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		metric.URIRequestTotal.With(prometheus.Labels{
			metric.Uri: ctx.Request.URL.RequestURI(),
			metric.Method: ctx.Request.Method,	
		}).Inc()
	}
}

func PrometheusUriErrorTotal() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.Request.URL.String() == "/metrics" {
			ctx.Next()
			return
		}
		ctx.Next()
		status := strconv.Itoa(ctx.Writer.Status())
		if status[0] != '2' {
			println("Error")
			metric.URIErrorTotal.With(prometheus.Labels{
				metric.Uri:			ctx.Request.URL.RequestURI(),
				metric.Method:		ctx.Request.Method,
				metric.StatusCode:	status,
			}).Inc()
		}
	}
}