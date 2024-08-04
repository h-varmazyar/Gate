package router

import (
	"github.com/gin-gonic/gin"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/controller/post"
	"go.uber.org/fx"
)

type V1 struct {
	postsController post.Controller
}

type V1Params struct {
	fx.In

	PostsController post.Controller
}

type V1Result struct {
	fx.Out

	V1Router *V1
}

func NewV1(params V1Params) V1Result {
	router := &V1{
		postsController: params.PostsController,
	}

	return V1Result{V1Router: router}
}

func (r *V1) RegisterRoutes(ginRouter *gin.RouterGroup) {
	v1Router := ginRouter.Group("/v1")

	{
		posts := v1Router.Group("/posts")
		posts.POST("/polarity", r.postsController.SubmitPolarity)
		posts.GET("/non-polarity", r.postsController.NonPolarityList)
	}
}
