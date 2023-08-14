package routes

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/kjpopov/gingonic-oauth-example/config"
)

func InitOauthCallbackRoute(r *gin.Engine) {
	// OAuth2 callback endpoint
	r.GET("/oauth2/callback", func(c *gin.Context) {
		code := c.Query("code")
		if code == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Missing code parameter"})
			return
		}

		token, err := config.Config.Oauth2Config.Exchange(context.Background(), code)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to exchange code for token"})
			return
		}

		// Serialize the token to JSON
		tokenJSON, err := json.Marshal(token)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to serialize token"})
			return
		}

		// Store the serialized token in the session
		session := sessions.Default(c)
		session.Set("oauth2_token", tokenJSON)
		session.Save()

		// Redirect to the / endpoint after successful authentication
		c.Redirect(http.StatusFound, "/")
	})
}
