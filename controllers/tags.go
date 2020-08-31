package controllers

import (
	"net/http"

	"task-api/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// GET /tags
// Get all tags
func FindTags(c *gin.Context) {

	userID, _ := c.Get("user_id")
	var tags []models.Tag
	models.DB.Where("user_id = ?", userID).Find(&tags)

	c.JSON(http.StatusOK, gin.H{"data": tags})
}

// POST /tags
// Create new tag
func CreateTag(c *gin.Context) {
	userID, _ := c.Get("user_id")

	// Validate input
	var tag models.Tag
	if err := c.ShouldBindJSON(&tag); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tag.UserID = uuid.MustParse(userID.(string))
	models.DB.Create(&tag)

	c.JSON(http.StatusOK, gin.H{"data": tag})
}

// GET /tags/:id
// Find a tag
func FindTag(c *gin.Context) { // Get model if exist
	userID, _ := c.Get("user_id")

	var tag models.Tag

	if err := models.DB.Where("id = ? AND user_id = ?", c.Param("id"), userID).First(&tag).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": tag})
}

// PUT /tags/:id
// Update a tag
func UpdateTag(c *gin.Context) {
	userID, _ := c.Get("user_id")

	// Get model if exist
	var tag models.Tag
	if err := models.DB.Where("id = ? AND user_id = ?", c.Param("id"), userID).First(&tag).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	if err := c.ShouldBindJSON(&tag); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tag.UserID = uuid.MustParse(userID.(string))
	models.DB.Model(&tag).Updates(tag)

	c.JSON(http.StatusOK, gin.H{"data": tag})
}

// DELETE /tags/:id
// Delete a tag
func DeleteTag(c *gin.Context) {
	userID, _ := c.Get("user_id")

	// Get model if exist
	var tag models.Tag
	if err := models.DB.Where("id = ? AND user_id = ?", c.Param("id"), userID).First(&tag).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&tag)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
