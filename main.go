package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	db "github.com/phillip-england/cfa-app/lib/db"
	pages "github.com/phillip-england/cfa-app/lib/pages"
	routes "github.com/phillip-england/cfa-app/lib/routes"
)

func main() {

	// configuration
	godotenv.Load()
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")

	// database
	client, err := db.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Disconnect(client)

	// pages
	pageSignup := pages.NewPageSignup(router)
	pageSignup.Mount()
	pageLogin := pages.NewPageLogin(router)
	pageLogin.Mount()
	pageLocation := pages.NewPageLocation(router)
	pageLocation.Mount()

	// api
	routeCreateUser := routes.NewRouteCreateUser(router, client)
	routeCreateUser.Mount()
	routeLoginUser := routes.NewRouteLoginUser(router, client)
	routeLoginUser.Mount()
	routeCreateLocation := routes.NewRouteCreateLocation(router, client)
	routeCreateLocation.Mount()




	router.Run()

}
