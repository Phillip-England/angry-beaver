package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PageLogin struct {
	router *gin.Engine
}

func NewPageLogin(router *gin.Engine) PageLogin {
	return PageLogin{
		router: router,
	}
}

func (r PageLogin) Mount() {
	r.router.GET("/", func(c *gin.Context) {
		r.Render(c)
	})
}

func (r PageLogin) Render(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
}

func (r PageLogin) Redirect(c *gin.Context) {
	c.Redirect(http.StatusUnauthorized, "/")
}