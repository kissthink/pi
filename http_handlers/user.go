package http_handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/smhouse/pi/db"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var form db.User_t
	if c.Bind(&form) == nil {
		
	}
}

func GetUser(c *gin.Context) {
	name := c.Param("name")
	u := db.User_t{Name: name}
	if err := u.Find(); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, u)
}
