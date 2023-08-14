package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/logger"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
	"github.com/kjpopov/gingonic-oauth-example/config"
	"github.com/kjpopov/gingonic-oauth-example/routes"
)

func main() {
	// Loading App Configuration
	config.LoadConfig(".")

	// Create a new Gin router
	r := gin.Default()
	store := memstore.NewStore([]byte("sup3r-$ecur3-s3cr37"))
	r.Use(sessions.Sessions("githubsession", store))

	// Add Logger middleware
	r.Use(logger.SetLogger())

	// Add Recovery middleware to handle panics
	r.Use(gin.Recovery())

	// Add CORS middleware to handle Cross-Origin requests
	r.Use(cors.Default())

	r.Static("/assets", "./assets")

	routes.InitHomeRoute(r)
	routes.InitCounterRoute(r)
	routes.InitOauthCallbackRoute(r)

	// Start the server
	r.Run(":8080")
}
