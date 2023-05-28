package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/phillip-england/cfa-app/lib/models/session"
	"github.com/phillip-england/cfa-app/lib/models/user"
	"go.mongodb.org/mongo-driver/mongo"
)

func Auth(c *gin.Context, client *mongo.Client) (*user.UserDoc, error) {
	
		// getting the session token
		token, err := c.Cookie("session")
		if err != nil {
			return nil, err
		}

		// creating a session db
		sessiondb := session.NewSessionDb(client, nil)

		// pulling the session document
		sessionDoc, err := sessiondb.FindById(token)
		if err != nil {
			return nil, err
		}

		// checking if the session token is expired
		if (time.Now().After(*sessionDoc.Expires)) {
			return nil, fmt.Errorf("expired session")
		}

		// creating a user db
		userdb := user.NewUserDb(client, nil)

		// pulling the user by the id
		userdoc, err := userdb.FindById(sessionDoc.User)
		if err != nil {
			return nil, err
		}

		return userdoc, nil
}