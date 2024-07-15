package Controller

import (
	"MyPlant-User/Database"
	"MyPlant-User/Model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var user Model.User
	var err error

	err = c.ShouldBindJSON(&user)
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	_, err = Model.GetUserByUsername(user.Username, Database.DB)
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
		return 
	}

	err = user.Save(Database.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created"})
}