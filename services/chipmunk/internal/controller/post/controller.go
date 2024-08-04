package post

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type Controller struct {
}

type Params struct {
	fx.In
}

type Result struct {
	fx.Out

	Controller *Controller
}

func New(p Params) Result {
	controller := &Controller{}
	return Result{Controller: controller}
}

func (c Controller) SubmitPolarity(ctx *gin.Context) {

}

func (c Controller) NonPolarityList(ctx *gin.Context) {

}
