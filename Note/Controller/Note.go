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
		Username: "",
		Link: "",
	}

	username , _ := c.Get("username")
	if err := c.ShouldBindJSON(&note); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	note.Username = username.(string)

	id, err := note.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to save note", "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": id.InsertedID,
		"username": note.Username,
		"title": note.Title,
		"content": note.Content,
		"status": note.Status,
	})
}

func UpdateNoteHandler(c *gin.Context) {
	username, _ := c.Get("username")
	id := c.Param("id")
	
	data, err := Model.FindOne(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get note", "error": err.Error()})
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
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update note", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Note updated",
		"result": result,
	})
}

func GetAllNoteHandler(c *gin.Context) {
	username, _ := c.Get("username")

	data, err := Model.Find(username.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get notes", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}

func GetNoteByIDHandler(c *gin.Context) {
	id := c.Param("id")
	username, _ := c.Get("username")

	note,err := Model.FindOne(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get note", "error": err.Error()})
		return
	}

	if note.Username != username.(string) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	c.JSON(http.StatusOK, note)
}

func DeleteNoteHandler(c *gin.Context) {
	username, _ := c.Get("username")
	id := c.Param("id")

	data, err := Model.FindOne(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get note", "error": err.Error()})
		return
	}

	if data.Username != username.(string) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	result, err := Model.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete note", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Note deleted",
		"result": result,
	})
}