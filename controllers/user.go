package controllers

import (
	"net/http"
	"task-api/auth"

	"task-api/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Hash function
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// VerifyPassword function
func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func Login(c *gin.Context) {

	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}

	var user models.User
	if err := models.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	if err := VerifyPassword(user.Password, input.Password); err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		c.JSON(http.StatusUnauthorized, "Please provide valid login details")
		return
	}

	token, err := auth.CreateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user, "token": token})
}

// GET /users
// Get all users
func FindUsers(c *gin.Context) {
	var users []models.User
	models.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}

// POST /users
// Create new user
func CreateUser(c *gin.Context) {
	// Validate input
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, _ := Hash(user.Password)
	user.Password = string(hashedPassword)

	models.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// GET /users/:id
// Find a user
func FindUser(c *gin.Context) { // Get model if exist
	var user models.User

	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

type RequestUpdateUserInput struct {
	User UpdateUserInput `json:"user" binding:"required"`
}
type UpdateUserInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// PUT /users/:id
// Update a user
func UpdateUser(c *gin.Context) {
	// Get model if exist
	var user models.User
	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input RequestUpdateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&user).Updates(input.User)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// DELETE /users/:id
// Delete a user
func DeleteUser(c *gin.Context) {
	// Get model if exist
	var user models.User
	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&user)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
