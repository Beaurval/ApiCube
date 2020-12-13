package controllers

import (
	"ApiCubes/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//FindTypeRessources Get all ressource
func FindTypeRessources(c *gin.Context) {
	var TypeRessource []models.TypeRessource
	models.DB.Find(&TypeRessource)

	c.JSON(http.StatusOK, gin.H{"data": TypeRessource})
}

//FindTypeRessource récupère la ressource correspondante à l'id passé en paramètre
func FindTypeRessource(c *gin.Context) {
	var TypeRessource models.TypeRessource

	if err := models.DB.Where("id = ?", c.Param("id")).First(&TypeRessource).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": TypeRessource})
}

// UpdateTypeRessource update les informations d'un citoyen
func UpdateTypeRessource(c *gin.Context) {
	// Get model if exist
	var TypeRessource models.TypeRessource
	if err := models.DB.Where("id = ?", c.Param("id")).First(&TypeRessource).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input models.UpdateTypeRessourceInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&TypeRessource).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": TypeRessource})
}

//CreateTypeRessource ajoute un type de ressource
func CreateTypeRessource(c *gin.Context) {
	//Validate input
	var input models.CreateTypeRessourceInput

	fmt.Println(c.Keys)
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create type ressource
	TypeRessource := models.TypeRessource{
		Nom: input.Nom,
	}
	models.DB.Create(&TypeRessource)

	c.JSON(http.StatusOK, gin.H{"data": TypeRessource})
}

// DeleteTypeRessource supprimer un type de ressource
func DeleteTypeRessource(c *gin.Context) {
	// Get model if exist
	var TypeRessource models.TypeRessource
	if err := models.DB.Where("id = ?", c.Param("id")).First(&TypeRessource).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&TypeRessource)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
