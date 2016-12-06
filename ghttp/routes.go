package ghttp

import "github.com/smhouse/pi/http_handlers"

func initRoutes() {
	v1 := router.Group("/v1")

	user := v1.Group("/user")
	user.POST("/", http_handlers.CreateUser)
	user.GET("/:name", http_handlers.GetUser)
}
