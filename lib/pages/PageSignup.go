package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PageSignup struct {
	router *gin.Engine
}

func NewPageSignup(router *gin.Engine) PageSignup {
	return PageSignup{
		router: router,
	}
}

func (r PageSignup) Mount() {
		r.router.GET("/signup", func(c *gin.Context) {
		c.HTML(http.StatusOK, "signup.html", nil)
	})
}