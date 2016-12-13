package http_handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/smhouse/pi/jwt"
)

func AuthJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Access-Token")
		if token == "" {
			c.JSON(http.StatusForbidden, error_t{Message: "Wrong access token"})
			c.Abort()
			return
		}

		usr, err := jwt.CheckToken(token)
		if err != nil {
			c.JSON(http.StatusForbidden, error_t{Message: err.Error()})
			c.Abort()
			return
		}

		c.Set("user", usr)
		c.Next()
	}
}