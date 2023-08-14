package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/kjpopov/gingonic-oauth-example/config"
	"golang.org/x/oauth2"
)

func InitHomeRoute(r *gin.Engine) {
	// Define the root route ("/") to check for authentication
	r.GET("/", func(c *gin.Context) {
		session := sessions.Default(c)
		// Get the OAuth2 token from session
		token := session.Get("oauth2_token")

		// TODO: Verify token
		if token == nil {
			// If the user is not authenticated, redirect to the /oauth2/callback endpoint for authentication
			c.Redirect(http.StatusFound, config.Config.Oauth2Config.AuthCodeURL("", oauth2.AccessTypeOffline))
			return
		}

		// If the user is authenticated, serve the index.html file
		c.File("index.html")
	})

	// Define the /hello endpoint and protect it with OAuth2
	r.GET("/hello", func(c *gin.Context) {
		// Retrieve the serialized token from the session
		session := sessions.Default(c)
		tokenJSON := session.Get("oauth2_token")
		if tokenJSON == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		// Deserialize the token from JSON
		var token oauth2.Token
		err := json.Unmarshal(tokenJSON.([]byte), &token)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to deserialize token"})
			return
		}

		// Get the authenticated user GitHub username:
		httpClient := config.Config.Oauth2Config.Client(context.Background(), &token)
		resp, err := httpClient.Get("https://api.github.com/user")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get github user username"})
			return
		}
		var respData map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&respData)
		if err != nil {
			// Handle the error
			return
		}

		// Extract the GitHub username from the response
		username, found := respData["login"].(string)
		if !found {
			c.JSON(200, gin.H{
				"message": "Hello, Anonymous User!",
			})
			return
		}

		c.JSON(200, gin.H{
			"message": fmt.Sprintf("Hello, %s!", username),
		})
	})
}
