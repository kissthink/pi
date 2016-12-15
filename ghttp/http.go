package ghttp

import (
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
	"fmt"
)

var router *gin.Engine

func StartHTTP(port int) error {
	router = gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})

	router.StaticFile("/favicon.ico", "./static/icon.ico")
	initRoutes()

	return router.Run(fmt.Sprintf(":%d", port))
}
