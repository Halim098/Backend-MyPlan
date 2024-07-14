package Router

import (
	"MyPlant-User/Controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/register", Controller.Register)
	r.POST("/login", Controller.Login)
	return r
}