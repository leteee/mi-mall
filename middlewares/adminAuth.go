package middlewares

import (
	"encoding/json"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"mi-mall/models"
	"strings"
)

func InitAdminAuthMiddleware(c *gin.Context) {
	//fmt.Println("InitAdminAuthMiddleware")
	//判断用户是否登录
	pathname := strings.Split(c.Request.URL.String(), "?")[0] //获取url路径去掉Get传值
	session := sessions.Default(c)
	userinfo := session.Get("userinfo")
	userinfoStr, ok := userinfo.(string)
	if ok {
		var u []models.Manager
		json.Unmarshal([]byte(userinfoStr), &u)
		if !(len(u) > 0 && u[0].Username != "") {
			if pathname != "/admin/login" && pathname != "/admin/doLogin" && pathname != "/admin/captcha" {
				c.Redirect(302, "/admin/login")
			}
		}
	} else {
		if pathname != "/admin/login" && pathname != "/admin/doLogin" && pathname != "/admin/captcha" {
			c.Redirect(302, "/admin/login")
		}
	}
	//fmt.Println(time.Now())
	//定义一个goroutine统计日志  当在中间件或 handler 中启动新的 goroutine 时，不能使用原始的上下文（c *gin.Context）， 必须使用其只读副本（c.Copy()）
	//cCp := c.Copy()
	//go func() {
	//	time.Sleep(2 * time.Second)
	//	fmt.Println("Done! in path " + cCp.Request.URL.Path)
	//}()
}
