package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phillip-england/cfa-app/lib/middleware"
	"github.com/phillip-england/cfa-app/lib/models/location"
	pages "github.com/phillip-england/cfa-app/lib/pages"
	"go.mongodb.org/mongo-driver/mongo"
)

type RouteCreateLocation struct {
	router *gin.Engine
	client *mongo.Client
}

func NewRouteCreateLocation(router *gin.Engine, client *mongo.Client) RouteCreateLocation {
	return RouteCreateLocation{
		router: router,
		client: client,
	}
}

func (r RouteCreateLocation) Mount() {

	// setting up our post route to take data from our form
	r.router.POST("/location", func(c *gin.Context) {
		
		// pulling in the log in page for rendering on failure to auth
		loginPage := pages.NewPageLogin(r.router)

		// authenticating the user
		user, err := middleware.Auth(c, r.client)
		if err != nil {
			loginPage.Render(c)
			return
		}

		// pulling in the location page which can render / refresh the page
		pageLocation := pages.NewPageLocation(r.router)

		// pulling the name and the number from the form submission
		name := c.PostForm("name")
		number := c.PostForm("number")

		// creating a location model from our user
		locationModel := location.NewLocationModel(user.ID, name, number)

		// validating our model
		err = locationModel.Validate()
		if err != nil {
			pageLocation.RenderWithError(c, err.Error())
			return
		}

		// using the location model to create a location db
		locationdb := location.NewLocationDb(r.client, &locationModel)

		// ensuring the location is unique
		err = locationdb.AssertUnique()
		if err != nil {
			pageLocation.RenderWithError(c, err.Error())
			return
		}

		// inserting the location
		_ = locationdb.Insert()

		// rendering the location page
		pageLocation.Render(c)

	})
}

// on exiting, we simply reload the location page
func (r RouteCreateLocation) Exit(c *gin.Context, err error) {
	c.HTML(http.StatusOK, "location.html", gin.H{
		"err_signup": err.Error(),
	})
}