package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"mi-mall/models"
	"mi-mall/routers"
)

func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()
	//自定义模版函数 注意要把这个函数放在加载模版前
	r.SetFuncMap(template.FuncMap{
		"UnixToTime": models.UnixToTime,
	})
	// 加载模版
	r.LoadHTMLGlob("templates/**/**/*")
	r.Static("/static", "static")

	routers.AdminRoutersInit(r)
	r.Run()
}
