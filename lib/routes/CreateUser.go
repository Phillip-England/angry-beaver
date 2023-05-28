package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/phillip-england/cfa-app/lib/models/user"
	"go.mongodb.org/mongo-driver/mongo"
)

type RouteCreateUser struct {
	router *gin.Engine
	client *mongo.Client
}

func NewRouteCreateUser(router *gin.Engine, client *mongo.Client) RouteCreateUser {
	return RouteCreateUser{
		router: router,
		client: client,
	}
}

func (r RouteCreateUser) Mount() {

	// setting up our post route to take data from our form
	r.router.POST("/signup", func(c *gin.Context) {

		// creating a user model
		model := user.NewUserModel(c.PostForm("email"), c.PostForm("password"))

		// validating the user
		err := model.Validate()
		if err != nil {
			r.Exit(c, err)
			return
		}

		// hashing the users password
		model.HashPassword()

		// creating a user database
		userdb := user.NewUserDb(r.client, &model)

		// ensuring the user is unique
		err = userdb.AssertUnique()
		if err != nil {
			r.Exit(c, err)
			return
		}

		// inserting the user
		_ = userdb.Insert()

		// redirecting to the login page if the user is created
		c.HTML(http.StatusOK, "index.html", nil)

	})
}

// on exiting, we simply reload the sign up page
func (r RouteCreateUser) Exit(c *gin.Context, err error) {
	c.HTML(http.StatusOK, "signup.html", gin.H{
		"err_signup": err.Error(),
	})
}