package server

import (
	"goprom/config"
	"goprom/delivery"
	"goprom/delivery/middleware"

	"github.com/gin-gonic/gin"
)

type AppServer interface {
	Run()
}

type appServer struct {
	routerEngine	*gin.Engine
	apiBaseUrl		string
	apiGroup		string
}

func (a *appServer) instrumentationMiddleware() {
	a.routerEngine.Use(middleware.PrometheusUriRequestTotal())
}

func (a *appServer) metricsHandler() {
	delivery.NewPrometheusApi(a.routerEngine)
}

func(a *appServer) handlers() {
	publicRoute := a.routerEngine.Group(a.apiGroup)
	delivery.NewStudentApi(publicRoute)
}

func (a *appServer) Run() {
	a.instrumentationMiddleware()
	a.metricsHandler()
	a.handlers()
	err := a.routerEngine.Run(a.apiBaseUrl)
	if err != nil {
		panic("Error run server")
	}
}

func NewAppServer() AppServer {
	newServer := new(appServer)

	c := config.NewConfig()

	r := gin.Default()
	newServer.routerEngine = r
	newServer.apiBaseUrl = c.ApiBaseUrl()
	newServer.apiGroup = c.ApiGroup()
	return newServer
}