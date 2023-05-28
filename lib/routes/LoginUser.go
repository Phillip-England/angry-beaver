package routes

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/phillip-england/cfa-app/lib/models/session"
	"github.com/phillip-england/cfa-app/lib/models/user"
	"github.com/phillip-england/cfa-app/lib/security"
	"go.mongodb.org/mongo-driver/mongo"
)

type RouteLoginUser struct {
	router *gin.Engine
	client *mongo.Client
}

func NewRouteLoginUser(router *gin.Engine, client *mongo.Client) RouteLoginUser {
	return RouteLoginUser{
		router: router,
		client: client,
	}
}

func (r RouteLoginUser) Mount() {

	// setting up our post route to take data from our form
	r.router.POST("/", func(c *gin.Context) {

		// creating a user model
		model := user.NewUserModel(c.PostForm("email"), c.PostForm("password"))

		// creating a user database
		userdb := user.NewUserDb(r.client, &model)

		// pulling the provided user
		userDoc, err := userdb.Pull()
		if err != nil {
			r.Exit(c, fmt.Errorf("invalid credentials"))
			return
		}

		// comparing the passwords
		err = security.Compare(userDoc.Password, model.Password)
		if err != nil {
			r.Exit(c, fmt.Errorf("invalid credentials"))
			return
		}

		// building a new user session
		sessionModel := session.NewSessionModel(userDoc.ID)

		// building a session db
		sessionDb := session.NewSessionDb(r.client, &sessionModel)

		// dropping all user sessions
		sessionDb.DropUserSessions()

		// inserting the session
		sessionDoc := sessionDb.Insert()
		
		// building a http only cookie containing the session _id
		now := time.Now()
		exp := now.Add(24 * time.Hour)
		cookie := http.Cookie{
			Name: "session",
			Value: sessionDoc.ID.Hex(),
			Expires: exp,
			HttpOnly: true,
		}
		http.SetCookie(c.Writer, &cookie)

		// if the user already has locations, redirect to the "Select a location page"

		// if they do not have any locations, send to the creation page
		c.Redirect(http.StatusFound, "/location")

	})
}

// on exiting, we simply reload the sign up page
func (r RouteLoginUser) Exit(c *gin.Context, err error) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"err_login": err.Error(),
	})
}