package controllers

import (
	"ApiCubes/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//FindCitoyens Get all citoyens
func FindCitoyens(c *gin.Context) {
	var citoyens []models.Citoyen
	models.DB.Preload("Rang").Find(&citoyens)

	c.JSON(http.StatusOK, gin.H{"data": citoyens})
}

//FindCitoyen récupère le citoyen correspondant à l'id passé en paramètre
func FindCitoyen(c *gin.Context) {
	var citoyen models.Citoyen

	if err := models.DB.Preload("Rang").Preload("Relations").Preload("Relations.TypeRelation").Preload("InRelations").Preload("InRelations.TypeRelation").Where("id = ?", c.Param("id")).First(&citoyen).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": citoyen})
}

// UpdateCitoyen update les informations d'un citoyen
func UpdateCitoyen(c *gin.Context) {
	// Get model if exist
	var citoyen models.Citoyen
	if err := models.DB.Where("id = ?", c.Param("id")).First(&citoyen).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input models.UpdateCitoyenInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&citoyen).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": citoyen})
}

//CreateCitoyen Add a citoyen
func CreateCitoyen(c *gin.Context) {
	//Validate input
	var input models.CreateCitoyenInput

	fmt.Println(c.Keys)
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create citoyen
	citoyen := models.Citoyen{
		Mail:       input.Mail,
		Ville:      input.Ville,
		CodePostal: input.CodePostal,
		Nom:        input.Nom,
		Prenom:     input.Prenom,
		MotDePasse: input.MotDePasse,
		Genre:      input.Genre,
		Adresse:    input.Adresse,
		Pseudo:     input.Pseudo,
		Telephone:  input.Telephone,
		RangID:     1,
	}
	models.DB.Create(&citoyen)

	c.JSON(http.StatusOK, gin.H{"data": citoyen})
}

// DeleteCitoyen supprimer un citoyen
func DeleteCitoyen(c *gin.Context) {
	// Get model if exist
	var citoyen models.Citoyen
	if err := models.DB.Where("id = ?", c.Param("id")).First(&citoyen).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&citoyen)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

// VoterRessource ajouter un vote à une ressource
func VoterRessource(c *gin.Context) {
	// Get model if exist
	var input models.RessourceVoted

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get ressource if exist
	var citoyen models.Citoyen
	if err := models.DB.Where("id = ?", input.CitoyenID).First(&citoyen).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Get ressource if exist
	var ressource models.Ressource
	if err := models.DB.Where("id = ?", input.RessourceID).First(&ressource).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Model(&citoyen).Association("RessourcesVoted").Append(&ressource)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

// VoterRessource ajouter un vote à une ressource
func ViewRessource(c *gin.Context) {
	// Get model if exist
	var input models.RessourceViewed

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get ressource if exist
	var citoyen models.Citoyen
	if err := models.DB.Where("id = ?", input.CitoyenID).First(&citoyen).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Get ressource if exist
	var ressource models.Ressource
	if err := models.DB.Where("id = ?", input.RessourceID).First(&ressource).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Model(&citoyen).Association("RessourcesViewed").Append(&ressource)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

// RetirerVoteRessource supprimer le vote d'une ressource
func RetirerVoteRessource(c *gin.Context) {
	// Get ressource if exist
	var citoyen models.Citoyen
	if err := models.DB.Where("id = ?", c.Param("idCitoyen")).First(&citoyen).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Get ressource if exist
	var ressource models.Ressource
	if err := models.DB.Where("id = ?", c.Param("idRessource")).First(&ressource).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Model(&citoyen).Association("RessourcesVoted").Delete(&ressource)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

// RetirerVoteCommentaire retirer le vote d'un commentaire
func RetirerVoteCommentaire(c *gin.Context) {

	var citoyen models.Citoyen
	if err := models.DB.Where("id = ?", c.Param("idCitoyen")).First(&citoyen).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var commentaire models.Commentaire
	if err := models.DB.Where("id = ?", c.Param("idCommentaire")).First(&commentaire).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Model(&citoyen).Association("CommentairesVoted").Delete(&commentaire)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

// VoterCommentaire ajouter un vote à une ressource
func VoterCommentaire(c *gin.Context) {
	// Get model if exist
	var input models.CommentaireVoted

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get commentaire if exist
	var commentaire models.Commentaire
	if err := models.DB.Where("id = ?", input.CommentaireID).First(&commentaire).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Get ressource if exist
	var citoyen models.Citoyen
	if err := models.DB.Where("id = ?", input.CitoyenID).First(&citoyen).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Model(&citoyen).Association("CommentairesVoted").Append(&commentaire)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
