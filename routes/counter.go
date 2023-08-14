package routes

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func InitCounterRoute(r *gin.Engine) {
	r.GET("/count-plus", func(c *gin.Context) {
		session := sessions.Default(c)
		var count int
		v := session.Get("count")
		if v == nil {
			count = 0
		} else {
			count = v.(int)
			count++
		}
		session.Set("count", count)
		session.Save()
		if count == 10 {
			c.Redirect(http.StatusFound, "/count-minus")
		}
		c.JSON(200, gin.H{"count": count})
	})

	r.GET("/count-minus", func(c *gin.Context) {
		session := sessions.Default(c)
		var count int
		v := session.Get("count")
		if v == nil {
			count = 0
		} else {
			count = v.(int)
			count--
		}
		session.Set("count", count)
		session.Save()
		if count == -10 {
			c.Redirect(http.StatusFound, "/count-plus")
		}
		c.JSON(200, gin.H{"count": count})
	})

}
