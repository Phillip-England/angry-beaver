package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PageLocation struct {
	router *gin.Engine
}


func NewPageLocation(router *gin.Engine) PageLocation {
	return PageLocation{
		router: router,
	}
}

// makes the route avilable to the client
func (r PageLocation) Mount() {
	r.router.GET("/location", func(c *gin.Context) {
		r.Render(c)
	})
}

// render function used for all standard page loads
func (r PageLocation) Render(c *gin.Context) {
	c.HTML(http.StatusOK, "location.html", gin.H{
		"banner": "Locations",
	})
}

// render function used to display the page with errors
func (r PageLocation) RenderWithError(c *gin.Context, formErr string) {
	c.HTML(http.StatusOK, "location.html", gin.H{
		"banner": "Locations",
		"formErr": formErr,
	})

}