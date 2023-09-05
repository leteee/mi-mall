package admin

import (
	"fmt"
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
		con.success(c, "登录成功", "/admin")
	} else {
		con.error(c, "验证码验证失败", "/admin/login")
	}
}
