package controllers

import (
	"CRUD/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateUserInput struct {
	Name   string `json:"name" binding:"required"`
	Family string `json:"family" binding:"required"`
	Age    uint   `json:"age" binding:"required"`
}

func CreateUser(c *gin.Context) {
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{Name: input.Name, Family: input.Family, Age: input.Age}
	models.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

type UpdateUserInput struct {
	Name   string `json:"name" binding:"required"`
	Family string `json:"family" binding:"required"`
	Age    uint   `json:"age" binding:"required"`
}

func UpdateUser(c *gin.Context) {
	var user models.User
	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "record not found"})
		return
	}

	var input UpdateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	updatedUser := models.User{Name: input.Name, Family: input.Family, Age: input.Age}
	models.DB.Model(&user).Updates(&updatedUser)

	c.JSON(http.StatusOK, gin.H{"data": user})

}

func DeleteUser(c *gin.Context) {
	var user models.User
	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "record not found"})
		return
	}

	models.DB.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"data": "successfuly deleted"})

}

func FindUsers(c *gin.Context) {
	var users []models.User
	models.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}
