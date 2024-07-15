package Router

import (
	"MyPlant-User/Controller"
	"MyPlant-User/Middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/register", Controller.Register)
	r.POST("/login", Controller.Login)

	r.Use(Middleware.Auth())
	r.GET("/user", Controller.GetUser)
	return r
}