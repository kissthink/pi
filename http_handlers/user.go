package http_handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/smhouse/pi/db"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var form db.User_t
	if c.Bind(&form) == nil {
		err := form.Find()
		if err == nil {
			c.JSON(http.StatusBadRequest, error_t{Message: "User already exists"})
			return
		}

		err = form.Create()
		if err != nil {
			c.JSON(http.StatusBadRequest, error_t{Message: err.Error()})
			return
		}

		c.JSON(http.StatusOK, form)
		return
	}

	c.JSON(http.StatusBadRequest, error_t{Message: "Validation error"})
}

func GetUser(c *gin.Context) {
	name := c.Param("name")
	u := db.User_t{Name: name}
	if err := u.Find(); err != nil {
		c.JSON(http.StatusBadRequest, error_t{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, u)
}
