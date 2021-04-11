package controllers

import (
	"ApiCubes/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindRelationsDuCitoyen(c *gin.Context) {
	var relationsCitoyen []models.RelationCitoyen

	models.DB.Preload("TypeRelation").Preload("Citoyen").Preload("CitoyenCible").Where("citoyen_id = ?", c.Param("id")).Find(&relationsCitoyen)

	c.JSON(http.StatusOK, gin.H{"data": relationsCitoyen})
}

func UpdateRelation(c *gin.Context) {
	// Get model if exist
	var relation models.RelationCitoyen

	// Validate input
	var input models.UpdateRelationCitoyenInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.DB.Where("citoyen_id = ?", input.CitoyenID).Where("citoyen_cible_id = ?", input.CitoyenCibleID).First(&relation).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	fmt.Printf("%v+\n", input)
	models.DB.Model(&relation).Where("citoyen_id = ? AND citoyen_cible_id = ?", input.CitoyenID, input.CitoyenCibleID).Update("Approbation", input.Approbation)
	c.JSON(http.StatusOK, gin.H{"data": relation})
}

func FindRelationsOuEstLeCitoyen(c *gin.Context) {
	var relationsCitoyen []models.RelationCitoyen

	models.DB.Preload("TypeRelation").Preload("Citoyen").Preload("CitoyenCible").Where("citoyen_cible_id = ?", c.Param("id")).Find(&relationsCitoyen)

	c.JSON(http.StatusOK, gin.H{"data": relationsCitoyen})
}

func AjouterRelation(c *gin.Context) {
	//Validate input
	var input models.CreateRelationCitoyenInput

	fmt.Println(c.Keys)
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create tag
	relationCitoyen := models.RelationCitoyen{
		CitoyenID:      input.CitoyenID,
		CitoyenCibleID: input.CitoyenCibleID,
		TypeRelationID: input.TypeRelationID,
		Approbation:    input.Approbation,
	}
	models.DB.Create(&relationCitoyen)

	c.JSON(http.StatusOK, gin.H{"data": relationCitoyen})
}

func DeleteRelation(c *gin.Context) {
	// Get model if exist
	var relation models.RelationCitoyen
	if err := models.DB.Where("citoyen_id = ?", c.Param("id")).First(&relation).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&relation)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
