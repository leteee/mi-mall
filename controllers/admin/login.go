package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginController struct {
	BaseController
}

func (con LoginController) Login(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/login/index.html", gin.H{})
}

func (con LoginController) DoLogin(c *gin.Context) {
	c.String(http.StatusOK, "编辑")
}
