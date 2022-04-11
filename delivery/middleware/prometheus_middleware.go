package middleware

import (
	"goprom/metric"
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