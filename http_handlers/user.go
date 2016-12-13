package http_handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/smhouse/pi/db"
	"net/http"
	"github.com/smhouse/pi/jwt"
)

func CreateUser(c *gin.Context) {
	//TODO: add email check
	var form user_create_form
	if c.Bind(&form) == nil {
		u := db.User_t{}
		u.Init(form.Name, form.Email, form.Password)

		err := u.Create()
		if err != nil {
			c.JSON(http.StatusBadRequest, error_t{Message: err.Error()})
			return
		}

		form.Password = ""
		c.JSON(http.StatusOK, form)
		return
	}

	c.JSON(http.StatusBadRequest, error_t{Message: "Validation error"})
}

func UpdateUser(c *gin.Context) {
	usr, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusForbidden, error_t{})
		return
	}
	user := usr.(*db.User_t)

	var form user_update_form
	if c.Bind(&form) == nil {
		user.Email = form.Email

		err := user.Update(form.Name, form.Password, form.NewPassword)
		if err != nil {
			c.JSON(http.StatusBadRequest, error_t{Message: err.Error()})
			return
		}

		c.JSON(http.StatusOK, user)
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

	u.Password = ""
	c.JSON(http.StatusOK, u)
}

func LoginUser(c *gin.Context) {
	var form user_login_form
	if c.Bind(&form) == nil {
		u := db.User_t{Name: form.Name}
		err := u.Find()
		if err != nil {
			c.JSON(http.StatusBadRequest, error_t{Message: err.Error()})
			return
		}

		err = u.ValidatePassword(form.Password)
		if err != nil {
			c.JSON(http.StatusBadRequest, error_t{Message: err.Error()})
			return
		}

		token, err := jwt.CreateToken(&u)
		if err != nil {
			c.JSON(http.StatusBadRequest, error_t{Message: err.Error()})
			return
		}

		c.JSON(http.StatusOK, map[string]string{
			"token": *token,
		})
		return
	}

	c.JSON(http.StatusBadRequest, error_t{Message: "Validation error"})
}