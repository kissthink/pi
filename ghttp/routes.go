package ghttp

import (
	"github.com/smhouse/pi/http_handlers"
	"github.com/gin-contrib/cors"
	"github.com/betacraft/yaag/yaag"
)

func initRoutes() {
	yaag.Init(&yaag.Config{On: true, DocTitle: "PI HUB", DocPath: "documentation.html"})
	router.Use(http_handlers.Document())
	router.Use(cors.Default())
	v1 := router.Group("/v1")

	v1.POST("/user", http_handlers.CreateUser)
	v1.GET("/user/:name", http_handlers.GetUser)
	v1.POST("/user/login", http_handlers.LoginUser)
	v1.PUT("/user", http_handlers.AuthJWT(), http_handlers.UpdateUser)

	v1.Use(http_handlers.AuthJWT())
	v1.POST("/device", http_handlers.CreateDevice)
	v1.PUT("/device/:id", http_handlers.UpdateDevice)
	v1.DELETE("/device/:id", http_handlers.DeleteDevice)
	v1.GET("/device/list", http_handlers.ListDevices)
}
