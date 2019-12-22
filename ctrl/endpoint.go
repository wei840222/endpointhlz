package ctrl

import (
	"github.com/kataras/iris/v12"
	"github.com/wei840222/endpointhlz/svc"
)

type EndpointCtrl struct{}

func NewEndpointCtrl() *EndpointCtrl {
	return &EndpointCtrl{}
}

func (ctrl *EndpointCtrl) InsertEndpoint(ctx iris.Context, endpointSVC *svc.EndpointSVC) {
	var e svc.Endpoint
	if err := ctx.ReadJSON(&e); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}
	if err := endpointSVC.InsertEndpoint(&e); err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(iris.Map{"message": "success"})
}

func (ctrl *EndpointCtrl) GetAllEndpoint(ctx iris.Context, endpointSVC *svc.EndpointSVC) {
	e, err := endpointSVC.GetAllEndpoint()
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString(err.Error())
		return
	}
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(e)
}
