package routers

import (
	"github.com/gin-gonic/gin"
	"mi-mall/controllers/itying"
)

func DefaultRoutersInit(r *gin.Engine) {
	defaultRouters := r.Group("/")
	{
		defaultRouters.GET("/", itying.DefaultController{}.Index)
		defaultRouters.GET("/news", itying.DefaultController{}.News)

	}
}
