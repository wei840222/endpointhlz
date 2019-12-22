package main

import (
	"github.com/wei840222/endpointhlz/ctrl"
	"github.com/wei840222/endpointhlz/svc"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/hero"

	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
)

func main() {
	app := iris.New()
	app.Use(recover.New())
	app.Use(logger.New())

	hero.Register(svc.NewEndpointSVC())

	app.Post("/api/endpoint", hero.Handler(ctrl.NewEndpointCtrl().InsertEndpoint))
	app.Get("/api/endpoint", hero.Handler(ctrl.NewEndpointCtrl().GetAllEndpoint))

	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
