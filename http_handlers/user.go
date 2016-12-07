package http_handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/smhouse/pi/db"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var form user_create_form
	if c.Bind(&form) == nil {
		u := db.User_t{}
		u.Init(form.Name, form.Email, form.Password)

		err := u.Create()
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

func LoginUser(c *gin.Context) {
	var form user_login_form
	if c.Bind(&form) == nil {

	}

	c.JSON(http.StatusBadRequest, error_t{Message: "Validation error"})
}