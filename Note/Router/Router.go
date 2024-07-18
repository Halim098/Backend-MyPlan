package Router

import (
	"MyPlan-Note/Controller"
	"MyPlan-Note/Middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(Middleware.Auth())

	r.GET("/note", Controller.GetAllNoteHandler)
	r.GET("/note/:id", Controller.GetNoteByIDHandler)
	r.POST("/note", Controller.InsertNoteHandler)
	r.PUT("/note/:id", Controller.UpdateNoteHandler)
	r.DELETE("/note/:id", Controller.DeleteNoteHandler)
	return r
}