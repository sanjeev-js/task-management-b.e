package controllers

import (
	"fmt"
	"net/http"
	"task-management/config"
	"task-management/models"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	fmt.Println(c, 'c')
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := validate.Struct(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userData := models.User{
		ID:        uuid.New().String(),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		Email:     user.Email,
		Phone:     user.Phone,
		Password:  user.Password,
		Role:      user.Role,
	}

	if err := config.DB.Create(&userData).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User Created successfully!"})

}
