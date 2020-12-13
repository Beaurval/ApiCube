package controllers

import (
	"ApiCubes/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//FindTypeRelations Get all relation
func FindTypeRelations(c *gin.Context) {
	var typeRelation []models.TypeRelation
	models.DB.Find(&typeRelation)

	c.JSON(http.StatusOK, gin.H{"data": typeRelation})
}

//FindTypeRelation récupère la relation correspondante à l'id passé en paramètre
func FindTypeRelation(c *gin.Context) {
	var typeRelation models.TypeRelation

	if err := models.DB.Where("id = ?", c.Param("id")).First(&typeRelation).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": typeRelation})
}

// UpdateTypeRelation update les informations d'un citoyen
func UpdateTypeRelation(c *gin.Context) {
	// Get model if exist
	var typeRelation models.TypeRelation
	if err := models.DB.Where("id = ?", c.Param("id")).First(&typeRelation).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input models.UpdateTypeRelationInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&typeRelation).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": typeRelation})
}

//CreateTypeRelation ajoute un type de relation
func CreateTypeRelation(c *gin.Context) {
	//Validate input
	var input models.CreateTypeRelationInput

	fmt.Println(c.Keys)
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create type relation
	typeRelation := models.TypeRelation{
		Nom: input.Nom,
	}
	models.DB.Create(&typeRelation)

	c.JSON(http.StatusOK, gin.H{"data": typeRelation})
}

// DeleteTypeRelation supprimer un type de relation
func DeleteTypeRelation(c *gin.Context) {
	// Get model if exist
	var typeRelation models.TypeRelation
	if err := models.DB.Where("id = ?", c.Param("id")).First(&typeRelation).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&typeRelation)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
