package ghttp

import "github.com/smhouse/pi/http_handlers"

func initRoutes() {
	v1 := router.Group("/v1")

	user := v1.Group("/user")
	user.POST("/", http_handlers.CreateUser)
	user.GET("/:name", http_handlers.GetUser)
	user.POST("/login", http_handlers.LoginUser)
	user.PUT("/", http_handlers.AuthJWT(), http_handlers.UpdateUser)

	device := v1.Group("/device")
	device.Use(http_handlers.AuthJWT())
	device.POST("/", http_handlers.CreateDevice)
	device.PUT("/:id", http_handlers.UpdateDevice)
	device.DELETE("/:id", http_handlers.DeleteDevice)
	device.GET("/list", http_handlers.ListDevices)

}
