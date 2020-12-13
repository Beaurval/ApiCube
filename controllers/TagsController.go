package controllers

import (
	"ApiCubes/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//FindTags Get all tags
func FindTags(c *gin.Context) {
	var Tag []models.Tag
	models.DB.Preload("Ressources").Find(&Tag)

	c.JSON(http.StatusOK, gin.H{"data": Tag})
}

//FindTag récupère le tag correspondant à l'id passé en paramètre
func FindTag(c *gin.Context) {
	var Tag models.Tag

	if err := models.DB.Preload("Ressources").Where("id = ?", c.Param("id")).First(&Tag).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Tag})
}

// UpdateTag update le tag
func UpdateTag(c *gin.Context) {
	// Get model if exist
	var Tag models.Tag
	if err := models.DB.Where("id = ?", c.Param("id")).First(&Tag).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input models.UpdateTagInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&Tag).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": Tag})
}

//CreateTag ajoute un tag
func CreateTag(c *gin.Context) {
	//Validate input
	var input models.CreateTagInput

	fmt.Println(c.Keys)
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create tag
	Tag := models.Tag{
		Nom: input.Nom,
	}
	models.DB.Create(&Tag)

	c.JSON(http.StatusOK, gin.H{"data": Tag})
}

// DeleteTag supprimer un tag
func DeleteTag(c *gin.Context) {
	// Get model if exist
	var Tag models.Tag
	if err := models.DB.Where("id = ?", c.Param("id")).First(&Tag).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&Tag)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
