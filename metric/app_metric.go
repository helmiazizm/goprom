package metric

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const (
	Uri			= "uri"
	Method		= "method"
	StatusCode 	= "statusCode"
	Stats		= "stats"
)

var (
	URIRequestTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "goprom_uri_request_total",
			Help: "all the server received request num with every uri",
		}, []string{Uri, Method},
	)

	URIErrorTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "goprom_uri_error_total",
			Help: "all the server error request with every uri",
		}, []string{Uri, Method, StatusCode},
	)
)