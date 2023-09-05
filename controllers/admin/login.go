package admin

import (
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"mi-mall/models"
	"net/http"
)

type LoginController struct {
	BaseController
}

func (con LoginController) Login(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/login/index.html", gin.H{})
}

func (con LoginController) Captcha(c *gin.Context) {
	id, b64s, err := models.MakeCaptcha()
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"captchaId":    id,
		"captchaImage": b64s,
	})
}

func (con LoginController) DoLogin(c *gin.Context) {
	captchaId := c.PostForm("captchaId")
	captchaText := c.PostForm("captchaText")
	fmt.Println(captchaText)
	if models.Verify(captchaId, captchaText) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		password = models.Md5(password)
		var userinfo []models.Manager
		models.DB.Where("username=? AND password=?", username, password).Find(&userinfo)
		fmt.Println(userinfo)
		if len(userinfo) > 0 {
			session := sessions.Default(c)
			userinfoSlice, _ := json.Marshal(userinfo)
			fmt.Println(string(userinfoSlice))
			session.Set("userinfo", string(userinfoSlice))
			session.Save()

			con.success(c, "登录成功", "/admin")
		} else {
			con.error(c, "用户名或密码错误", "/admin/login")
		}
	} else {
		con.error(c, "验证码验证失败", "/admin/login")
	}
}

func (con LoginController) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("userinfo")
	session.Save()
	con.success(c, "退出登录成功", "/admin/login")
}
