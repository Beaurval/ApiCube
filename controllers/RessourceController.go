package controllers

import (
	"ApiCubes/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//FindRessources Récupérer toutes les ressources
func FindRessources(c *gin.Context) {
	var ressources []models.Ressource
	models.DB.Find(&ressources)

	c.JSON(http.StatusOK, gin.H{"data": ressources})
}

//FindRessource récupère la ressource correspondant à l'id passé en paramètre
func FindRessource(c *gin.Context) {
	var ressource models.Ressource

	if err := models.DB.Where("id = ?", c.Param("id")).First(&ressource).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": ressource})
}

//CreateRessource Ajoute une nouvelle ressource
func CreateRessource(c *gin.Context) {
	//Validate input
	var input models.CreateRessourceInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create Ressource
	ressource := models.Ressource{
		Titre:          input.Titre,
		Contenu:        input.Contenu,
		CitoyenID:      input.CitoyenID,
		TypeRelationID: input.TypeRelationID,
		Vues:           0,
		Votes:          0,
	}
	models.DB.Create(&ressource)

	c.JSON(http.StatusOK, gin.H{"data": ressource})
}

// UpdateRessource update les informations d'une ressource
func UpdateRessource(c *gin.Context) {
	// Get model if exist
	var ressource models.Ressource
	if err := models.DB.Where("id = ?", c.Param("id")).First(&ressource).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input models.UpdateRessourceInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&ressource).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": ressource})
}

// DeleteRessource supprimer une ressource
func DeleteRessource(c *gin.Context) {
	// Get model if exist
	var ressource models.Ressource
	if err := models.DB.Where("id = ?", c.Param("id")).First(&ressource).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&ressource)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
