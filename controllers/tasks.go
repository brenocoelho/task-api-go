package controllers

import (
	"net/http"

	"task-api/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// GET /tasks
// Get all tasks
func FindTasks(c *gin.Context) {

	userID, _ := c.Get("user_id")
	var tasks []models.Task
	models.DB.Where("user_id = ?", userID).Find(&tasks)

	c.JSON(http.StatusOK, gin.H{"data": tasks})
}

// POST /tasks
// Create new task
func CreateTask(c *gin.Context) {
	userID, _ := c.Get("user_id")

	// Validate input
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task.UserID = uuid.MustParse(userID.(string))
	models.DB.Create(&task)

	c.JSON(http.StatusOK, gin.H{"data": task})
}

// GET /tasks/:id
// Find a task
func FindTask(c *gin.Context) { // Get model if exist
	userID, _ := c.Get("user_id")

	var task models.Task

	if err := models.DB.Where("id = ? AND user_id = ?", c.Param("id"), userID).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": task})
}

// PUT /tasks/:id
// Update a task
func UpdateTask(c *gin.Context) {
	userID, _ := c.Get("user_id")

	// Get model if exist
	var task models.Task
	if err := models.DB.Where("id = ? AND user_id = ?", c.Param("id"), userID).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task.UserID = uuid.MustParse(userID.(string))
	models.DB.Model(&task).Updates(task)

	c.JSON(http.StatusOK, gin.H{"data": task})
}

// DELETE /tasks/:id
// Delete a task
func DeleteTask(c *gin.Context) {
	userID, _ := c.Get("user_id")

	// Get model if exist
	var task models.Task
	if err := models.DB.Where("id = ? AND user_id = ?", c.Param("id"), userID).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&task)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
