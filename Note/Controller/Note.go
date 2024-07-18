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
