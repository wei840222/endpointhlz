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

	e := svc.NewEndpointSVC()
	svc.NewHlzSVC(e).StartCronJob()
	hero.Register(e)

	app.HandleDir("/", "vue/dist")

	api := app.Party("/api")
	{
		api.Post("/endpoint", hero.Handler(ctrl.NewEndpointCtrl().InsertEndpoint))
		api.Get("/endpoint", hero.Handler(ctrl.NewEndpointCtrl().GetAllEndpoint))
		api.Put("/endpoint/{id:int64}", hero.Handler(ctrl.NewEndpointCtrl().UpdateEndpointByID))
		api.Delete("/endpoint/{id:int64}", hero.Handler(ctrl.NewEndpointCtrl().DeleteEndpointByID))
	}

	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}
