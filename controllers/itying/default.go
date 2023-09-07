package itying

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type DefaultController struct{}

func (con DefaultController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/itying/index.html", gin.H{})

}
