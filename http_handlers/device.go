package http_handlers

import (
	"gopkg.in/gin-gonic/gin.v1"
	"github.com/smhouse/pi/db"
	"net/http"
	"strconv"
)

func CreateDevice(c *gin.Context) {
	usr, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusForbidden, error_t{})
		return
	}
	user := usr.(*db.User_t)

	var form db.Device_t
	if c.Bind(&form) == nil {
		form.UserName = user.Name
		err := form.Create()
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

func UpdateDevice(c *gin.Context) {
	usr, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusForbidden, error_t{})
		return
	}
	user := usr.(*db.User_t)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, error_t{Message: "Wrong id"})
		return
	}

	d := db.Device_t{ID: uint64(id)}
	err = d.Find()
	if err != nil {
		c.JSON(http.StatusBadRequest, error_t{Message: err.Error()})
		return
	}

	//TODO: maybe add old password validation
	if d.UserName != user.Name {
		c.JSON(http.StatusBadRequest, error_t{Message: "Wrong user"})
		return
	}

	if c.Bind(&d) == nil {
		err := d.Update()
		if err != nil {
			c.JSON(http.StatusBadRequest, error_t{Message: err.Error()})
			return
		}

		d.Password = ""
		c.JSON(http.StatusOK, d)
		return
	}

	c.JSON(http.StatusBadRequest, error_t{Message: "Validation error"})
}

func DeleteDevice(c *gin.Context) {
	usr, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusForbidden, error_t{})
		return
	}
	user := usr.(*db.User_t)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, error_t{Message: "Wrong id"})
		return
	}

	d := db.Device_t{ID: uint64(id)}
	err = d.Find()
	if err != nil {
		c.JSON(http.StatusBadRequest, error_t{Message: err.Error()})
		return
	}

	if d.UserName != user.Name {
		c.JSON(http.StatusBadRequest, error_t{Message: "Wrong user"})
		return
	}

	err = d.Delete()
	if err != nil {
		c.JSON(http.StatusBadRequest, error_t{Message: err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func ListDevices(c *gin.Context) {
	usr, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusForbidden, error_t{Message: "No user"})
		return
	}
	user := usr.(*db.User_t)

	d := db.Device_t{UserName: user.Name}
	devices, err := d.List()
	if err != nil {
		c.JSON(http.StatusBadRequest, error_t{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, devices)
}
