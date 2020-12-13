package controllers

import (
	"ApiCubes/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//FindRangs Get all relation
func FindRangs(c *gin.Context) {
	var rang []models.Rang
	models.DB.Find(&rang)

	c.JSON(http.StatusOK, gin.H{"data": rang})
}

//FindRang récupère la relation correspondante à l'id passé en paramètre
func FindRang(c *gin.Context) {
	var rang models.Rang

	if err := models.DB.Where("id = ?", c.Param("id")).First(&rang).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": rang})
}

// UpdateRang update les informations d'un citoyen
func UpdateRang(c *gin.Context) {
	// Get model if exist
	var rang models.Rang
	if err := models.DB.Where("id = ?", c.Param("id")).First(&rang).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input models.UpdateRangInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&rang).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": rang})
}

//CreateRang ajoute un type de relation
func CreateRang(c *gin.Context) {
	//Validate input
	var input models.CreateRangInput

	fmt.Println(c.Keys)
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create type relation
	rang := models.Rang{
		Nom: input.Nom,
	}
	models.DB.Create(&rang)

	c.JSON(http.StatusOK, gin.H{"data": rang})
}

// DeleteRang supprimer un type de relation
func DeleteRang(c *gin.Context) {
	// Get model if exist
	var rang models.Rang
	if err := models.DB.Where("id = ?", c.Param("id")).First(&rang).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&rang)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
