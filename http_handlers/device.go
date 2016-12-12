package http_handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/smhouse/pi/db"
	"net/http"
	"log"
)

func CreateDevice(c *gin.Context) {
	usr, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusForbidden, error_t{})
		return
	}

	log.Printf("%+v", usr)

	var form db.Device_t
	if c.Bind(&form) == nil {
		//form.UserName = user.Name
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
