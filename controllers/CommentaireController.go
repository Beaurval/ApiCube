package controllers

import (
	"ApiCubes/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//FindCommentaires Récupérer touts les commentaires
func FindCommentaires(c *gin.Context) {
	var commentaire []models.Commentaire
	models.DB.Preload("Reponses").Preload("Citoyen").Preload("CitoyenVoted").Find(&commentaire)

	c.JSON(http.StatusOK, gin.H{"data": commentaire})
}

//FindCommentaire récupère la commentaire correspondant à l'id passé en paramètre
func FindCommentaire(c *gin.Context) {
	var commentaire models.Commentaire

	if err := models.DB.Preload("Reponses").Preload("Citoyen").Preload("CitoyenVoted").Where("id = ?", c.Param("id")).First(&commentaire).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": commentaire})
}

//CreateCommentaire Ajoute un nouveau commentaire
func CreateCommentaire(c *gin.Context) {
	//Validate input
	var input models.CreateCommentaireInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create commentaire
	commentaire := models.Commentaire{
		Contenu:     input.Contenu,
		RessourceID: input.RessourceID,
		ParentID:    input.ParentID,
		CitoyenID:   input.CitoyenID,
	}
	models.DB.Create(&commentaire)

	c.JSON(http.StatusOK, gin.H{"data": commentaire})
}

// UpdateCommentaire update les informations d'un commentaire
func UpdateCommentaire(c *gin.Context) {
	// Get model if exist
	var commentaire models.Commentaire
	if err := models.DB.Preload("Reponses").Where("id = ?", c.Param("id")).First(&commentaire).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		fmt.Println(err)
		return
	}

	// Validate input
	var input models.UpdateCommentaireInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	models.DB.Model(&commentaire).Updates(&models.Commentaire{
		Contenu: input.Contenu,
		Vote:    input.Vote,
	})

	c.JSON(http.StatusOK, gin.H{"data": commentaire})
}

// DeleteCommentaire supprimer un commentaire
func DeleteCommentaire(c *gin.Context) {
	// Get model if exist
	var commentaire models.Commentaire
	if err := models.DB.Where("id = ?", c.Param("id")).First(&commentaire).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Select("Reponses").Delete(&commentaire)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
