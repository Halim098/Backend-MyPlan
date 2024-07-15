package Controller

import (
	"MyPlan-User/Database"
	"MyPlan-User/Helper"
	"MyPlan-User/Model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userInput struct {
	Username string `json:"username"`
	Password string `json:"password" binding:"required"`
	NewPassword string `json:"new_password"`
}

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

func Login(c *gin.Context) {
	var user Model.User
	var err error

	err = c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	dbUser,err := Model.GetUserByUsername(user.Username, Database.DB)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	err = dbUser.ValidatePassword(user.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong Password"})
		return
	}

	token, err := Helper.GenerateToken(dbUser.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token":token})
}

func GetUser(c *gin.Context) {
	username, _ := c.Get("username")

	user, err := Model.GetUserByUsername(username.(string), Database.DB)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": user.ID, "username": user.Username, "created_at" : user.CreatedAt, "updated_at" : user.UpdatedAt})
}

func UpdateUser(c *gin.Context) {
	username, _ := c.Get("username")

	var user *Model.User

	var userInput userInput = userInput{
		Username: "",
		Password: "",
		NewPassword: "",
	}

	var err error

	err = c.ShouldBindJSON(&userInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if userInput.Password == ""  || (userInput.Username == "" && userInput.NewPassword == "") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	user, err = Model.GetUserByUsername(username.(string), Database.DB)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	err = user.ValidatePassword(userInput.Password)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Wrong Password"})
		fmt.Println(userInput.Password)
		return
	}

	if userInput.Username != "" {
		err = user.UpdateUsername(userInput.Username, Database.DB)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update username"})
			return
		}

		user,err = Model.GetUserByUsername(userInput.Username, Database.DB)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Username updated", "username": user.Username})
	}

	if userInput.NewPassword != "" {
		user.Password = userInput.NewPassword
		user.BeforeSave(Database.DB)
		err = user.UpdatePassword(user.Password, Database.DB)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update password"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Password updated"})
	}
}