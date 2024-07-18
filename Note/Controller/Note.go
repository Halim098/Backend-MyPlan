package Controller

import (
	"MyPlan-Note/Model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InsertNoteHandler(c *gin.Context) {
	note := Model.Note{
		Title: "",
		Content: nil,
		Status: "private",
	}

	username , _ := c.Get("username")
	if err := c.ShouldBindJSON(&note); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	note.Username = username.(string)

	_, err := note.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Insert note"})
		return
	}

	c.JSON(http.StatusCreated, note)
}

func UpdateNoteHandler(c *gin.Context) {
	username, _ := c.Get("username")
	id := c.Param("id")
	
	data, err := Model.FindOne(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ID not found"})
		return
	}

	if data.Username != username.(string) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	result, err := data.Update(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update note"})
		return
	}

	c.JSON(http.StatusOK, result)
}

func GetAllNoteHandler(c *gin.Context) {
	username, _ := c.Get("username")

	data, err := Model.Find(username.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get notes"})
		return
	}

	c.JSON(http.StatusOK, data)
}
