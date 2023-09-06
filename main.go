package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"html/template"
	"mi-mall/models"
	"mi-mall/routers"
)

func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()
	// 创建基于cookie的存储引擎,secretKey参数是用于加密的密钥
	store := cookie.NewStore([]byte("secretKey"))
	r.Use(sessions.Sessions("mySession", store))
	// 设置session中间件，参数mySession指的是session的名字，也是cookie的名字
	// store是前面创建的存储引擎，我们可以替换成其他存储引擎
	r.Use(sessions.Sessions("mySession", store))
	//自定义模版函数 注意要把这个函数放在加载模版前
	r.SetFuncMap(template.FuncMap{
		"UnixToTime": models.UnixToTime,
	})
	// 加载模版
	r.LoadHTMLGlob("templates/**/**/*")
	r.Static("/static", "static")

	routers.AdminRoutersInit(r)
	r.Run(":80")
}
